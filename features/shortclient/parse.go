package shortclient

import (
	"strings"
)

const (
	lineSep = "\n"
)

func (htmlClient HTMLClient) parse(content string) []string {
	linesResult := make([]string, 0, 10)

	lines := strings.Split(content, lineSep)

	for _, line := range lines {
		for from, to := range htmlClient.replacements {
			updatedLine := strings.TrimSpace(strings.ReplaceAll(line, from, to))
			if updatedLine != "" {
				linesResult = append(linesResult, updatedLine)
			}
		}
	}

	return linesResult
}
