package bot

import (
	"strings"

	"gopkg.in/telebot.v4"
)

func (bot Bot) cmdURLsSetup() {
	tb := bot.telebot

	var inTxtHandler = func(ctx telebot.Context) error {
		var (
			message       = ctx.Message()
			caption       = message.Caption
			text          = message.Text
			actualCaption = strings.TrimSpace(caption)
			actualText    = strings.TrimSpace(text)
			userText      = ""
		)
		if actualText != "" {
			userText = actualText
		} else if actualCaption != "" {
			userText = actualCaption
		}

		urls := extract(userText)

		if len(urls) == 0 {
			return bot.respondZero(ctx)
		}

		if len(urls) == 1 {
			first := urls[0]
			return bot.respondSingle(ctx, first)
		}

		return bot.respondMulti(ctx, urls)
	}

	tb.Handle(telebot.OnForward, inTxtHandler)

	tb.Handle(telebot.OnText, inTxtHandler)
}
