package bot

import "gopkg.in/telebot.v4"

func (bot Bot) respondMulti(ctx telebot.Context, urls []string) error {
	inlineKeys := createInlineKeys(urls)
	return ctx.Send("Выберите ссылку:", &telebot.ReplyMarkup{
		InlineKeyboard: inlineKeys,
	})
}
