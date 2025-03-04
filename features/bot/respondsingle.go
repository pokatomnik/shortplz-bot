package bot

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
)

func (bot Bot) respondSingle(ctx telebot.Context, url string) error {
	var (
		users  = bot.users
		sc     = bot.summaryClient
		userId = ctx.Sender().ID
		user   = users.GetUser(userId)
	)
	if user.IsError() {
		return ctx.Send(errorFailedGetUser)
	}

	apiToken := user.MustGet().APIToken
	if apiToken == "" {
		return ctx.Send(messageNoToken)
	}

	shortInfo := sc.Get(apiToken, url)
	if shortInfo.IsError() {
		warnMsg := fmt.Sprintf("Failed to get summary for url: %s, error: %v", url, shortInfo.Error().Error())
		logrus.Warn(warnMsg)
		return ctx.Send(errorSummarizationFailed, &telebot.SendOptions{
			ReplyTo: ctx.Message(),
		})
	}

	linesJoined := strings.Join(shortInfo.MustGet(), "\n")
	return ctx.Send(linesJoined, &telebot.SendOptions{
		ReplyTo: ctx.Message(),
	})
}
