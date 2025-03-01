package bot

import (
	"github.com/pokatomnik/shortplz-bot/features/summary"
	"github.com/samber/mo"
	"gopkg.in/telebot.v4"
)

type Bot struct {
	telebot       *telebot.Bot
	summaryClient summary.SummaryClient
	apiToken      string
}

func New(token string, summaryClient summary.SummaryClient, apiToken string) mo.Result[Bot] {
	botPrefs := telebot.Settings{Token: token}
	telebot, err := telebot.NewBot(botPrefs)
	if err != nil {
		return mo.Err[Bot](err)
	}
	bot := Bot{telebot: telebot, summaryClient: summaryClient, apiToken: apiToken}

	return mo.Ok(bot)
}
