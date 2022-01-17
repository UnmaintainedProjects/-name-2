package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/bot/manager"
	"github.com/gotgcalls/tgcalls"
)

func resume(b *gotgbot.Bot, ctx *ext.Context) error {
	instance, ok, err := manager.CurrentManager.GetInstance(ctx.EffectiveChat.Id)
	if !ok {
		if err != nil {
			return err
		}
		return nil
	}
	result, err := instance.Resume()
	if err != nil {
		return err
	}
	switch result {
	case tgcalls.Ok:
		ctx.EffectiveMessage.Reply(b, "Resumed.", nil)
	case tgcalls.NotPaused:
		ctx.EffectiveMessage.Reply(b, "Not paused to resume.", nil)
	case tgcalls.NotInCall:
		ctx.EffectiveMessage.Reply(b, "Not in call.", nil)
	}
	return nil
}

var resumeHandler = handlers.NewCommand("resume", resume)
