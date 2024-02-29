package format

import "regexp"

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