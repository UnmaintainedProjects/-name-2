package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/bot/manager"
	"github.com/gotgcalls/tgcalls"
)

func pause(b *gotgbot.Bot, ctx *ext.Context) error {
	instance, ok, err := manager.CurrentManager.GetInstance(ctx.EffectiveChat.Id)
	if !ok {
		if err != nil {
			return err
		}
		return nil
	}
	result, err := instance.Pause()
	if err != nil {
		return err
	}
	switch result {
	case tgcalls.Ok:
		ctx.EffectiveMessage.Reply(b, "Paused.", nil)
	case tgcalls.NotStreaming:
		ctx.EffectiveMessage.Reply(b, "Not streaming to pause.", nil)
	case tgcalls.NotInCall:
		ctx.EffectiveMessage.Reply(b, "Not in call.", nil)
	}
	return nil
}

var pauseHandler = handlers.NewCommand("pause", pause)
