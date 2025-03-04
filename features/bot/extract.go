package bot

import (
	"fmt"
	"strings"

	"github.com/mvdan/xurls"
	"gopkg.in/telebot.v4"
)

const (
	https = "https://"
	http  = "http://"
)

func normalizeUrl(url string) string {
	isOk := strings.HasPrefix(url, https) || strings.HasPrefix(url, http)
	if isOk {
		return url
	}
	return fmt.Sprintf("%s%s", https, url)
}

func fromText(text string) []string {
	urls := xurls.Relaxed.FindAllString(text, -1)
	if urls == nil {
		// FindAllString may return nil when no matches
		urls = []string{}
	}
	for i := range urls {
		urls[i] = normalizeUrl(urls[i])
	}
	return urls
}

func fromEnts(ents telebot.Entities) []string {
	res := make([]string, 0, len(ents))
	for _, e := range ents {
		url := strings.TrimSpace(e.URL)
		if url != "" {
			res = append(res, normalizeUrl(e.URL))
		}
	}
	return res
}

func extractUrlsFromMessage(msg *telebot.Message) []string {
	var (
		text = strings.TrimSpace(msg.Text)
		capt = strings.TrimSpace(msg.Caption)
		ents = msg.CaptionEntities
	)

	turls := fromText(text)
	curls := fromText(capt)
	eurls := fromEnts(ents)

	res := make([]string, 0, len(turls)+len(curls)+len(eurls))

	added := make(map[string]struct{})

	for _, u := range turls {
		if _, ok := added[u]; !ok {
			res = append(res, u)
			added[u] = struct{}{}
		}
	}
	for _, u := range curls {
		if _, ok := added[u]; !ok {
			res = append(res, u)
			added[u] = struct{}{}
		}
	}
	for _, u := range eurls {
		if _, ok := added[u]; !ok {
			res = append(res, u)
			added[u] = struct{}{}
		}
	}

	return res
}
