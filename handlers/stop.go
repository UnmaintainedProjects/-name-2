package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/bot/manager"
)

func stop(b *gotgbot.Bot, ctx *ext.Context) error {
	instance, ok, err := manager.CurrentManager.GetInstance(ctx.EffectiveChat.Id)
	if !ok {
		if err != nil {
			return err
		}
		return nil
	}
	ok, err = instance.Stop()
	if err != nil {
		return err
	}
	if ok {
		ctx.EffectiveMessage.Reply(b, "Stopped.", nil)
	} else {
		ctx.EffectiveMessage.Reply(b, "Not in call.", nil)
	}
	return nil
}

var stopHandler = handlers.NewCommand("stop", stop)
