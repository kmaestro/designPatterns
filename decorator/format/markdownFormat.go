package format

import (
	"regexp"
	"strings"
)

type markdownFormat struct {
	next InputFormat
}

func NewMarkdownFormat(input InputFormat) markdownFormat {
	return markdownFormat{next: input}
}

func (m markdownFormat) FormatText(text string) string {
	chunks := strings.Split(text, "\n\n")
	for i, chunk := range chunks {
		if matched, _ := regexp.MatchString("^#+", chunk); matched {
			re := regexp.MustCompile("^(#+)(.*?)$")
			chunks[i] = re.ReplaceAllStringFunc(chunk, func(s string) string {
				matches := re.FindStringSubmatch(s)
				h := len(matches[1])
				return "<h" + string('0'+h) + ">" + strings.TrimSpace(matches[2]) + "</h" + string('0'+h) + ">"
			})
		} else {
			chunks[i] = "<p>" + chunk + "</p>"
		}
	}
	text = strings.Join(chunks, "\n\n")

	reStrongDouble := regexp.MustCompile(`\*\*(.*?)\*\*`)
	text = reStrongDouble.ReplaceAllString(text, "<strong>$1</strong>")
	reStrongUnderscore := regexp.MustCompile(`__(.*?)__`)
	text = reStrongUnderscore.ReplaceAllString(text, "<strong>$1</strong>")
	reEmAsterisk := regexp.MustCompile(`\*(.*?)\*`)
	text = reEmAsterisk.ReplaceAllString(text, "<em>$1</em>")
	reEmUnderscore := regexp.MustCompile(`_(.*?)_`)
	text = reEmUnderscore.ReplaceAllString(text, "<em>$1</em>")

	return text
}
