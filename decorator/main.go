package main

import (
	"fmt"
	"regexp"
)

type InputFormat interface {
	FormatText(text string) string
}

type dangerousHTMLTagsFilter struct {
	next InputFormat
}

func NewDangerousHTMLTagsFilter(input InputFormat) dangerousHTMLTagsFilter {
	return dangerousHTMLTagsFilter{next: input}
}

func (d dangerousHTMLTagsFilter) FormatText(text string) string {
	text = d.next.FormatText(text)

	re := regexp.MustCompile(`<script.*?>([\s\S]*)?</script>`)
	return re.ReplaceAllString(text, "")
}

type markdownFormat struct {
	next InputFormat
}

func NewMarkdownFormat(input InputFormat) markdownFormat {
	return markdownFormat{next: input}
}

func (m markdownFormat) FormatText(text string) string {
	text = m.next.FormatText(text)

	re := regexp.MustCompile(`\*\*(.*?)\*\*`)
	return re.ReplaceAllString(text, "<strong>$1</strong>")
}

type textInput struct {
}

func NewTextInput() *textInput {
	return &textInput{}
}

func (t textInput) FormatText(text string) string {
	return text
}

func main() {
	filter := NewDangerousHTMLTagsFilter(NewMarkdownFormat(NewTextInput()))

	text := "# Welcome\n"
	text += "This is my first post on this **gorgeous** forum.\n"
	text += "<script src=\"http://www.iwillhackyou.com/script.js\">\n"
	text += "performXSSAttack();\n"
	text += "</script>\n"
	fmt.Println(filter.FormatText(text))
}
