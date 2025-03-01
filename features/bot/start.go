package bot

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
)

func (bot Bot) Start() {
	var (
		tb    = bot.telebot
		sc    = bot.summaryClient
		token = bot.apiToken
	)

	tb.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send(messageWelcome)
	})

	tb.Handle("/help", func(ctx telebot.Context) error {
		return ctx.Send(helpStr)
	})

	tb.Handle(telebot.OnText, func(ctx telebot.Context) error {
		maybeURLStr := ctx.Message().Text

		_, err := url.ParseRequestURI(maybeURLStr)
		if err != nil {
			return ctx.Send(errorNotAnURL, &telebot.SendOptions{
				ReplyTo: ctx.Message(),
			})
		}

		shortInfo := sc.Get(token, maybeURLStr)
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
