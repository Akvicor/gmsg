package gmsg

import "strings"

var Telegram = newTelegram()

type telegram struct {
	shouldEscaped string
}

func newTelegram() *telegram {
	return &telegram{
		shouldEscaped: "_*[]()~`>#+-=|{}!.",
	}
}

func (t *telegram) Escape(str string) string {
	result := &strings.Builder{}
	for _, c := range str {
		if strings.ContainsRune(t.shouldEscaped, c) {
			result.WriteRune('\\')
		}
		result.WriteRune(c)
	}
	return result.String()
}
