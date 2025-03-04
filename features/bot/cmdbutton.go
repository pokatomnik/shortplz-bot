package bot

import (
	"gopkg.in/telebot.v4"
)

func (bot Bot) cmdButtonSetup() {
	bot.telebot.Handle(createEmptyURLButton(), func(ctx telebot.Context) error {
		url := ctx.Data()
		return bot.respondSingle(ctx, url)
	})
}
