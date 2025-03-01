package main

import (
	"fmt"

	"github.com/pokatomnik/shortplz-bot/features/bot"
	"github.com/pokatomnik/shortplz-bot/features/env"
	"github.com/pokatomnik/shortplz-bot/features/summary"
	"github.com/sirupsen/logrus"
)

const (
	errorBotTokenMissing    = "Telegram bot token is missing, can't proceed"
	errorAPITokenMissing    = "API token is missing"
	errorBotStartFailedBase = "Failed to start bot with reason"
	infoBotStarted          = "Bot started"
)

func main() {
	env := env.New()
	botToken := env.GetBotToken()
	apiToken := env.GetAPIToken()

	if botToken.IsAbsent() {
		logrus.Error(errorBotTokenMissing)
		return
	}
	if apiToken.IsAbsent() {
		logrus.Error(errorAPITokenMissing)
		return
	}

	summaryClient := summary.New()

	bot := bot.New(botToken.MustGet(), summaryClient, apiToken.MustGet())

	if bot.IsError() {
		logrus.Error(fmt.Sprintf("%s: %v", errorBotStartFailedBase, bot.Error()))
		return
	}

	logrus.Info(infoBotStarted)
	bot.MustGet().Start()
}
