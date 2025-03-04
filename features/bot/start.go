package bot

func (bot Bot) Start() {
	bot.cmdStartSetup()
	bot.cmdHelpSetup()
	bot.cmdTokenSetup()
	bot.cmdURLsSetup()
	bot.cmdButtonSetup()

	bot.telebot.Start()
}
