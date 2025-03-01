package bot

import (
	"fmt"
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
		return ctx.Send("Hello!")
	})

	tb.Handle(telebot.OnText, func(ctx telebot.Context) error {
		url := ctx.Message().Text
		shortInfo := sc.Get(token, url)
		if shortInfo.IsError() {
			logrus.Warn(fmt.Sprintf("Failed to get summary for url: %s, error: %v", url, shortInfo.Error().Error()))
			return ctx.Send(errorSummarizationFailed, &telebot.SendOptions{
				ReplyTo: ctx.Message(),
			})
		} else {
			linesJoined := strings.Join(shortInfo.MustGet(), "\n")
			return ctx.Send(linesJoined, &telebot.SendOptions{
				ReplyTo: ctx.Message(),
			})
		}
	})

	tb.Start()
}
