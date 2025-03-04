package bot

func (bot Bot) Start() {
	bot.cmdStartSetup()
	bot.cmdHelpSetup()
	bot.cmdTokenSetup()
	bot.cmdURLsSetup()

	bot.telebot.Start()
}
