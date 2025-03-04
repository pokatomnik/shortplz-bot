package bot

import (
	"fmt"
	"strings"

	"github.com/mvdan/xurls"
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

func extract(text string) []string {
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
