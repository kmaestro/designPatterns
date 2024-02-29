package format

type InputFormat interface {
	FormatText(text string) string
}