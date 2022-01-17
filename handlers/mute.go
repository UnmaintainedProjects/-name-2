package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/bot/manager"
	"github.com/gotgcalls/tgcalls"
)

func mute(b *gotgbot.Bot, ctx *ext.Context) error {
	instance, ok, err := manager.CurrentManager.GetInstance(ctx.EffectiveChat.Id)
	if !ok {
		if err != nil {
			return err
		}
		return nil
	}
	result, err := instance.Mute()
	if err != nil {
		return err
	}
	switch result {
	case tgcalls.Ok:
		ctx.EffectiveMessage.Reply(b, "Muted.", nil)
	case tgcalls.AlreadyMuted:
		ctx.EffectiveMessage.Reply(b, "Already muted.", nil)
	case tgcalls.NotInCall:
		ctx.EffectiveMessage.Reply(b, "Not in call.", nil)
	}
	return nil
}

var muteHandler = handlers.NewCommand("mute", mute)
