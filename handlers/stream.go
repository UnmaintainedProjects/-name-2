package handlers

import (
	"fmt"

	"github.com/gotgcalls/bot/downloader"
	"github.com/gotgcalls/bot/queues"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/bot/converter"
	"github.com/gotgcalls/bot/manager"
)

func stream(b *gotgbot.Bot, ctx *ext.Context) error {
	repliedMessage := ctx.EffectiveMessage.ReplyToMessage
	if repliedMessage == nil || repliedMessage.Audio == nil {
		ctx.EffectiveMessage.Reply(b, "Reply an audio file.", nil)
		return nil
	}
	msg, err := ctx.EffectiveMessage.Reply(b, "Downloading...", nil)
	if err != nil {
		return err
	}
	input, err := downloader.Download(b, repliedMessage.Audio)
	if err != nil {
		msg.EditText(b, "An error occurred while downloading.", nil)
		return nil
	}
	msg, _, err = msg.EditText(b, "Converting...", nil)
	if err != nil {
		return err
	}
	file, err := converter.Convert(input)
	if err != nil {
		msg.EditText(b, "An error occurred while converting.", nil)
		return nil
	}
	// if isFinished, _ := tgcalls.Get().IsFinished(tgcalls.CLIENT, ctx.EffectiveChat.Id); isFinished != gotgcalls.OK {
	// 	err = tgcalls.Get().Stream(tgcalls.CLIENT, ctx.EffectiveChat.Id, filePath)
	// 	if err != nil {
	// 		edit(b, msg, i18n.Localize("stream_error", map[string]string{"Error": err.Error()}))
	// 		return nil
	// 	}

	// 	edit(b, msg, i18n.Localize("streaming", nil))
	// } else {
	// 	position := queues.Push(ctx.EffectiveChat.Id, filePath)
	// 	edit(b, msg, i18n.Localize("queued_at", map[string]int{"Position": position}))
	// }
	position := queues.Push(ctx.EffectiveChat.Id, file)
	if position == 1 {
		instance, ok, err := manager.CurrentManager.GetInstance(ctx.EffectiveChat.Id)
		if !ok {
			if err != nil {
				return err
			}
			return nil
		}
		err = instance.Stream(file)
		if err != nil {
			msg.EditText(b, "An error occurred while attempting to start streaming.", nil)
			return nil
		}
		msg.EditText(b, "Streaming...", nil)
		return nil
	}
	msg.EditText(b, fmt.Sprintf("Queued at position %d.", position), nil)
	return nil
}

var streamHandler = handlers.NewCommand("stream", stream)
