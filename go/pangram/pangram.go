package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(input string) bool {

	if len(input) < 26 {
		return false
	}

	occurrences := make(map[int32]int)

	for _, c := range strings.ToLower(input) {
		occurrences[c] += 1
	}

	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		if occurrences[c] == 0 {
			return false
		}
	}

	return true
}
