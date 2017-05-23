package acronym

import (
	"strings"
)

const testVersion = 2

func Abbreviate(input string) string {
	result := ""

	noDashes := strings.Replace(input, "-", " ", -1)

	words := strings.Split(noDashes, " ")

	for _, word := range words {

		if len(word) > 0 {
			c := string(word[0])
			result += strings.ToUpper(c)
		}

		if word != strings.ToUpper(word) {
			for _, c := range word[1:] {
				if c >= 'A' && c <= 'Z' {
					result += string(c)
				}
			}
		}
	}

	return result
}
