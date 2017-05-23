package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(input string) bool {

	if len(input) < 26 {
		return false
	}

	s := strings.ToLower(input)

	for i := 'a'; i <= 'z'; i++ {
		if !strings.Contains(s, string(i)) {
			return false
		}
	}

	return true
}
