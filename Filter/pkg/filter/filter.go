package filter

import (
	"strings"
)

var filter = []string{
	"qwerty",
	"йцукен",
	"zxvbnm",
}

// Censorship возвращает false если в тексте найдены недопустимые слова
func Censorship(str string) bool {

	for _, s := range filter {
		if strings.Contains(str, s) {
			return true
		}
	}

	return false
}
