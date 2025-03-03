package bot

import (
	"github.com/pokatomnik/shortplz-bot/entities/user"
	"github.com/samber/mo"
	"gopkg.in/telebot.v4"
)

type SummaryClient interface {
	Get(token string, url string) mo.Result[[]string]
}

type UsersRepository interface {
	GetUser(userId int64) mo.Result[user.User]
	UpdateToken(userId int64, token string) mo.Result[struct{}]
}

type Bot struct {
	telebot       *telebot.Bot
	summaryClient SummaryClient
	users         UsersRepository
}

func New(token string, summaryClient SummaryClient, users UsersRepository) mo.Result[Bot] {
	botPrefs := telebot.Settings{Token: token}
	telebot, err := telebot.NewBot(botPrefs)
	if err != nil {
		return mo.Err[Bot](err)
	}
	bot := Bot{telebot: telebot, summaryClient: summaryClient, users: users}

	return mo.Ok(bot)
}
