package bot

import "gopkg.in/telebot.v4"

func (bot Bot) cmdStartSetup() {
	bot.telebot.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send(messageWelcome)
	})
}
