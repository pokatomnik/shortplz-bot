package bot

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/pokatomnik/shortplz-bot/entities/user"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
)

func (bot Bot) Start() {
	var (
		tb    = bot.telebot
		sc    = bot.summaryClient
		users = bot.users
	)

	tb.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send(messageWelcome)
	})

	tb.Handle("/help", func(ctx telebot.Context) error {
		return ctx.Send(helpStr)
	})

	tb.Handle("/token", func(ctx telebot.Context) error {
		text := ctx.Message().Text
		userId := ctx.Sender().ID
		words := strings.Fields(text)

		if len(words) < 2 {
			token := users.GetUser(userId).OrElse(user.User{}).APIToken
			if token == "" {
				token = emptyTokenCaption
			}

			return ctx.Send(fmt.Sprintf("%s%s", messageDisplayTokenPrefix, token), &telebot.SendOptions{
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

	tb.Handle(telebot.OnText, func(ctx telebot.Context) error {
		maybeURLStr := ctx.Message().Text

		_, err := url.ParseRequestURI(maybeURLStr)
		if err != nil {
			return ctx.Send(errorNotAnURL, &telebot.SendOptions{
				ReplyTo: ctx.Message(),
			})
		}

		userId := ctx.Sender().ID
		user := users.GetUser(userId)
		if user.IsError() {
			return ctx.Send(errorFailedGetUser)
		}

		apiToken := user.MustGet().APIToken
		if apiToken == "" {
			return ctx.Send(messageNoToken)
		}

		shortInfo := sc.Get(apiToken, maybeURLStr)
		if shortInfo.IsError() {
			logrus.Warn(fmt.Sprintf("Failed to get summary for url: %s, error: %v", maybeURLStr, shortInfo.Error().Error()))
			return ctx.Send(errorSummarizationFailed, &telebot.SendOptions{
				ReplyTo: ctx.Message(),
			})
		}

		linesJoined := strings.Join(shortInfo.MustGet(), "\n")
		return ctx.Send(linesJoined, &telebot.SendOptions{
			ReplyTo: ctx.Message(),
		})
	})

	tb.Start()
}
