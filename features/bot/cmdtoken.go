package bot

import (
	"fmt"
	"strings"

	"github.com/pokatomnik/shortplz-bot/entities/user"
	"gopkg.in/telebot.v4"
)

func (bot Bot) cmdTokenSetup() {
	var (
		tb    = bot.telebot
		users = bot.users
	)

	tb.Handle("/token", func(ctx telebot.Context) error {
		text := ctx.Message().Text
		userId := ctx.Sender().ID
		words := strings.Fields(text)

		if len(words) < 2 {
			token := users.GetUser(userId).OrElse(user.User{}).APIToken
			if token == "" {
				token = emptyTokenCaption
			}

			return ctx.Send(fmt.Sprintf("%s\"%s\"", messageDisplayTokenPrefix, token), &telebot.SendOptions{
				ReplyTo: ctx.Message(),
			})
		}

		res := users.UpdateToken(ctx.Sender().ID, words[1])
		if res.IsError() {
			return ctx.Send(messageFailedUpdateToken, &telebot.SendOptions{
				ReplyTo: ctx.Message(),
			})
		}

		return ctx.Send(messageTokenSaved, &telebot.SendOptions{
			ReplyTo: ctx.Message(),
		})
	})
}
