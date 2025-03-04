package bot

import "gopkg.in/telebot.v4"

const uid = "URL"

func createURLButton(text, url string) *telebot.InlineButton {
	return &telebot.InlineButton{Unique: uid, Text: text, Data: url}
}

func createEmptyURLButton() *telebot.InlineButton {
	return &telebot.InlineButton{Unique: uid}
}

func createInlineKeys(urls []string) [][]telebot.InlineButton {
	inlineKeys := make([][]telebot.InlineButton, 0, len(urls))
	for _, url := range urls {
		inlineKeys = append(inlineKeys, []telebot.InlineButton{
			*createURLButton(url, url),
		})
	}
	return inlineKeys
}
