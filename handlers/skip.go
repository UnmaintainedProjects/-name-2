package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/bot/manager"
	"github.com/gotgcalls/bot/queues"
	"github.com/gotgcalls/tgcalls"
)

func skip(b *gotgbot.Bot, ctx *ext.Context) error {
	instance, ok, err := manager.CurrentManager.GetInstance(ctx.EffectiveChat.Id)
	if !ok {
		if err != nil {
			return err
		}
		return nil
	}
	result, err := instance.Finish()
	if err != nil {
		return err
	}
	switch result {
	case tgcalls.Ok:
		queues.Skip(ctx.EffectiveChat.Id)
		ctx.EffectiveMessage.Reply(b, "Skipped.", nil)
	case tgcalls.NotStreaming:
		ctx.EffectiveMessage.Reply(b, "Nothing to skip.", nil)
	case tgcalls.NotInCall:
		ctx.EffectiveMessage.Reply(b, "Not in call.", nil)
	}
	return nil
}

var skipHandler = handlers.NewCommand("skip", skip)
