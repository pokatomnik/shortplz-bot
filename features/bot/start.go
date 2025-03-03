package bot

import (
	"fmt"
	"strings"

	"github.com/mvdan/xurls"
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

	tb.Handle(telebot.OnText, func(ctx telebot.Context) error {
		text := ctx.Message().Text

		urls := xurls.Strict.FindAllString(text, -1)
		if urls == nil {
			// FindAllString may return nil when no matches
			urls = []string{}
		}

		if len(urls) == 0 {
			return ctx.Send(errorNotAnURL, &telebot.SendOptions{
				ReplyTo: ctx.Message(),
			})
		}

		if len(urls) == 1 {
			first := urls[0]
			userId := ctx.Sender().ID
			user := users.GetUser(userId)
			if user.IsError() {
				return ctx.Send(errorFailedGetUser)
			}

			apiToken := user.MustGet().APIToken
			if apiToken == "" {
				return ctx.Send(messageNoToken)
			}

			shortInfo := sc.Get(apiToken, first)
			if shortInfo.IsError() {
				logrus.Warn(fmt.Sprintf("Failed to get summary for url: %s, error: %v", first, shortInfo.Error().Error()))
				return ctx.Send(errorSummarizationFailed, &telebot.SendOptions{
					ReplyTo: ctx.Message(),
				})
			}

			linesJoined := strings.Join(shortInfo.MustGet(), "\n")
			return ctx.Send(linesJoined, &telebot.SendOptions{
				ReplyTo: ctx.Message(),
			})
		}

		inlineKeys := make([][]telebot.InlineButton, 0, len(urls))
		for _, url := range urls {
			inlineKeys = append(inlineKeys, []telebot.InlineButton{
				{Unique: "URL", Text: url, Data: url},
			})
		}

		return ctx.Send("Выберите ссылку:", &telebot.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		})
		// return ctx.Send("Выберите ссылку:", &telebot.SendOptions{
		// 	ReplyParams: {
		// 		inlineKeys: inlineKeys,
		// 	},
		// 	ReplyTo: ctx.Message(),
		// })
	})

	tb.Handle(&telebot.InlineButton{Unique: "URL"}, func(ctx telebot.Context) error {
		data := ctx.Data()
		fmt.Println(data)
		return ctx.Send("Button pressed")
	})

	tb.Start()
}
