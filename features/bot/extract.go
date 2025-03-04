package bot

import "github.com/mvdan/xurls"

func extract(text string) []string {
	urls := xurls.Strict.FindAllString(text, -1)
	if urls == nil {
		// FindAllString may return nil when no matches
		urls = []string{}
	}
	return urls
}
