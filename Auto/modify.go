package Auto

import (
	"strings"
	"fmt"
	"strconv"
)

func Modify(text string) string {
	words := strings.Fields(text)
	fmt.Println(words)

	for i := 0; i < len(words); i++ {
		// word := words[i]

		switch words[i] {
		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(cap)":
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(up,":
			num, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err != nil {
				fmt.Println(err)
				break
			}
			for j := 0; j <= num; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		case "(low,":
			num, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err != nil {
				fmt.Println(err)
				break
			}
			for j := 0; j <= num; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		case "(cap,":
			num, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err != nil {
				fmt.Println(err)
				break
			}
			for j := 0; j <= num; j++ {
				words[i-j] = strings.Title(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		}
		text = strings.Join(words, " ")
		text = ReplaceBin(text)
	}
	return text
}
