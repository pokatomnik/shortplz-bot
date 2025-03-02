package bot

import (
	"github.com/pokatomnik/shortplz-bot/features/summary"
	userspkg "github.com/pokatomnik/shortplz-bot/features/users"
	"github.com/samber/mo"
	"gopkg.in/telebot.v4"
)

type Bot struct {
	telebot       *telebot.Bot
	summaryClient summary.SummaryClient
	users         userspkg.Users
}

func New(token string, summaryClient summary.SummaryClient, users userspkg.Users) mo.Result[Bot] {
	botPrefs := telebot.Settings{Token: token}
	telebot, err := telebot.NewBot(botPrefs)
	if err != nil {
		return mo.Err[Bot](err)
	}
	bot := Bot{telebot: telebot, summaryClient: summaryClient, users: users}

	return mo.Ok(bot)
}
