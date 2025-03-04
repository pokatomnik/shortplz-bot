package main

import (
	"fmt"

	"github.com/pokatomnik/shortplz-bot/features/apiclient"
	"github.com/pokatomnik/shortplz-bot/features/bot"
	"github.com/pokatomnik/shortplz-bot/features/env"
	"github.com/pokatomnik/shortplz-bot/features/shortclient"
	"github.com/pokatomnik/shortplz-bot/features/summary"
	"github.com/pokatomnik/shortplz-bot/features/users"
	"github.com/sirupsen/logrus"
)

const (
	errorBotTokenMissing    = "Telegram bot token is missing, can't proceed"
	errorAPITokenMissing    = "API token is missing"
	errorBotStartFailedBase = "Failed to start bot with reason"
	errorFailedOpenDB       = "Failed to open database"
	infoBotStarted          = "Bot started"
)

func main() {
	env := env.New()
	botToken := env.GetBotToken()
	dbFName := env.GetDBFileName()

	if botToken.IsAbsent() {
		logrus.Error(errorBotTokenMissing)
		return
	}

	apiClient := apiclient.New()
	shortClient := shortclient.New()
	summaryClient := summary.New(apiClient, shortClient)

	users := users.Open(dbFName.OrElse("data.db"))
	if users.IsError() {
		logrus.Error(errorFailedOpenDB)
		return
	}
	defer users.MustGet().Close()

	bot := bot.New(botToken.MustGet(), summaryClient, users.MustGet())

	if bot.IsError() {
		logrus.Error(fmt.Sprintf("%s: %v", errorBotStartFailedBase, bot.Error().Error()))
		return
	}

	logrus.Info(infoBotStarted)
	bot.MustGet().Start()
}
