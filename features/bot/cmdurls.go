package bot

import (
	"errors"

	"gopkg.in/telebot.v4"
)

func (bot Bot) cmdURLsSetup() {
	tb := bot.telebot

	var inTxtHandler = func(ctx telebot.Context) error {
		urls := extractUrlsFromMessage(ctx.Message())

		if len(urls) == 0 {
			return bot.respondZero(ctx)
		}

		first := urls[0]
		return bot.respondSingle(ctx, first)
	}

	// Some additional checks are required here.
	// ctx.Message().IsForwarded() should be invoked
	// to let us know if the message was forwarded
	// to enable one of two handlers,
	// because sometimes both are triggered
	tb.Handle(telebot.OnForward, func(ctx telebot.Context) error {
		if ctx.Message().IsForwarded() {
			return inTxtHandler(ctx)
		}
		return errors.New("telebot bug")
	})

	tb.Handle(telebot.OnText, func(ctx telebot.Context) error {
		if ctx.Message().IsForwarded() {
			return errors.New("telebot bug")
		}
		return inTxtHandler(ctx)
	})
}
