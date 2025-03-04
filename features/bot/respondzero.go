package bot

import "gopkg.in/telebot.v4"

func (bot Bot) respondZero(ctx telebot.Context) error {
	return ctx.Send(errorNotAnURL, &telebot.SendOptions{
		ReplyTo: ctx.Message(),
	})
}
