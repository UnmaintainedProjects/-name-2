package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func Add(dp *ext.Dispatcher) {
	handlers2 := []ext.Handler{
		streamHandler,
		muteHandler,
		unmuteHandler,
		pauseHandler,
		resumeHandler,
		skipHandler,
		stopHandler,
		endHandler,
	}
	dp.AddHandler(
		handlers.NewMessage(
			func(msg *gotgbot.Message) bool {
				return msg.Chat.Type == "supergroup"
			},
			func(b *gotgbot.Bot, ctx *ext.Context) error {
				for _, handler := range handlers2 {
					if handler.CheckUpdate(b, ctx.Update) {
						return handler.HandleUpdate(b, ctx)
					}
				}
				return nil
			},
		),
	)
}
