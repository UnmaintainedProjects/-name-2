package main

import (
	"context"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	auth2 "github.com/gotgcalls/bot/auth"
	"github.com/gotgcalls/bot/queues"
	"github.com/gotgcalls/tgcalls"
	"github.com/joho/godotenv"

	"github.com/gotgcalls/bot/handlers"
	"github.com/gotgcalls/bot/manager"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	client, err := telegram.ClientFromEnvironment(telegram.Options{
		SessionStorage: &session.FileStorage{Path: "session.json"},
		UpdateHandler: telegram.UpdateHandlerFunc(
			func(ctx context.Context, u tg.UpdatesClass) error {
				return manager.CurrentManager.Handle(ctx, u)
			},
		),
	})
	if err != nil {
		panic(err)
	}
	bot, err := gotgbot.NewBot(os.Getenv("BOT_TOKEN"), nil)
	if err != nil {
		panic(err)
	}
	updater := ext.NewUpdater(nil)
	handlers.Add(updater.Dispatcher)
	updater.StartPolling(bot, &ext.PollingOpts{DropPendingUpdates: true})
	client.Run(context.Background(), func(ctx context.Context) error {
		flow := auth.NewFlow(
			auth2.TermAuth{},
			auth.SendCodeOptions{},
		)
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return err
		}
		manager.CurrentManager = manager.New(ctx, client.API())
		manager.CurrentManager.OnFinish = func(chatId int64, instance *tgcalls.TGCalls) {
			i := queues.Pull(chatId)
			file, ok := i.(string)
			if !ok {
				return
			}
			instance.Stream(file)
		}
		return telegram.RunUntilCanceled(ctx, client)
	})
}
