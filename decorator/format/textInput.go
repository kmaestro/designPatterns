package format

type textInput struct {
}

func NewTextInput() *textInput {
	return &textInput{}
}

func (t textInput) FormatText(text string) string {
	return text
}