package main

import (
	"designPatterns/decorator/format"
	"fmt"
)

func main() {

	text := format.NewTextInput()
	markdown := format.NewMarkdownFormat(text)
	filteredInput := format.NewDangerousHTMLTagsFilter(markdown)
	fmt.Println(filteredInput.FormatText(`# Welcome


	This is my first post on this **gorgeous** forum.
	<script src=\"http://www.iwillhackyou.com/script.js\">
	performXSSAttack();
	</script>
	`))
}
