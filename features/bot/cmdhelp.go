package bot

import "gopkg.in/telebot.v4"

func (bot Bot) cmdHelpSetup() {
	bot.telebot.Handle("/help", func(ctx telebot.Context) error {
		return ctx.Send(helpStr)
	})
}
