package Auto

import (
	"strings"
	"fmt"
)

func Modify(text string) string {
	words := strings.Fields(text)
	fmt.Println(words)

	for i := 0; i < len(words); i++ {
		word := words[i]

		switch word {
		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
