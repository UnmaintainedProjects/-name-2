package handlers

import (
	"github.com/gotgcalls/bot/manager"
	"github.com/gotgcalls/bot/queues"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func end(b *gotgbot.Bot, ctx *ext.Context) error {
	queues.Clear(ctx.EffectiveChat.Id)
	manager.CurrentManager.TerminateInstance(ctx.EffectiveChat.Id)
	return nil
}

var endHandler = handlers.NewMessage(
	func(msg *gotgbot.Message) bool {
		return msg.VoiceChatEnded != nil
	},
	end,
)
