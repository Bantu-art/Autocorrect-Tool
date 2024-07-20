
package Auto

import (
	"strconv"
	"strings"
)
// Function to replace a binary number with its decimal equivalent
func ReplaceBin(text string) string {
	words := strings.Fields(text)
	for i, ch := range words {
		if ch == "(bin)" && i > 0 {
			// Convert the previous word to decimal
			decValue, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(decValue, 10)
			}
			// Remove the (bin) from the words
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}