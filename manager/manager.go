package manager

import (
	"context"
	"sync"

	"github.com/gotd/td/tg"
	"github.com/gotgcalls/bot/utils"
	"github.com/gotgcalls/tgcalls"
)

type Manager struct {
	ctx context.Context
	api *tg.Client

	instances    *sync.Map
	accessHashes *sync.Map

	OnFinish func(chatId int64, instance *tgcalls.TGCalls)
}

var CurrentManager *Manager

func New(ctx context.Context, api *tg.Client) *Manager {
	return &Manager{ctx: ctx, api: api, instances: &sync.Map{}, accessHashes: &sync.Map{}}
}

func (m *Manager) GetInstance(chatId int64) (*tgcalls.TGCalls, bool, error) {
	channelId := utils.ToMTProto(chatId)
	instance, ok := m.instances.Load(channelId)
	if !ok {
		accessHash, ok := m.accessHashes.Load(channelId)
		if !ok {
			return nil, false, nil
		}
		instance := tgcalls.New(m.ctx, &tg.InputChannel{ChannelID: channelId, AccessHash: accessHash.(int64)}, m.api, nil)
		instance.OnFinish = func() {
			if m.OnFinish != nil {
				m.OnFinish(utils.ToBotAPI(channelId), instance)
			}
		}
		err := tgcalls.Start(instance)
		if err != nil {
			return nil, false, err
		}
		m.instances.Store(channelId, instance)
		return instance, true, nil
	}
	return instance.(*tgcalls.TGCalls), true, nil
}

func (m *Manager) TerminateInstance(chatId int64) bool {
	channelId := utils.ToMTProto(chatId)
	i, ok := m.instances.Load(channelId)
	if !ok {
		return false
	}
	instance := i.(*tgcalls.TGCalls)
	instance.Stop()
	tgcalls.Stop(instance)
	m.instances.Delete(channelId)
	return true
}

func (m *Manager) Handle(ctx context.Context, u tg.UpdatesClass) error {
	switch u := u.(type) {
	case (*tg.Updates):
		if len(u.Chats) > 0 {
			for _, chat := range u.Chats {
				if chat, ok := chat.(*tg.Channel); ok {
					m.accessHashes.Store(chat.ID, chat.AccessHash)
				}
			}
		}
	}
	return nil
}
