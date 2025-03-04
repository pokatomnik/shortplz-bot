package bot

import (
	"gopkg.in/telebot.v4"
)

func (bot Bot) cmdURLsSetup() {
	tb := bot.telebot

	tb.Handle(telebot.OnText, func(ctx telebot.Context) error {
		text := ctx.Message().Text

		urls := extract(text)

		if len(urls) == 0 {
			return bot.respondZero(ctx)
		}

		if len(urls) == 1 {
			first := urls[0]
			bot.respondSingle(ctx, first)
		}

		return bot.respondMulti(ctx, urls)
	})
}
