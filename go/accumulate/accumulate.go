// Package accumulate takes an array of strings and applies a dynamically defined function
package accumulate

const testVersion = 1

func Accumulate(input []string, f func(string) string) []string {
	var newOutput []string
	for _, v := range input {
		newOutput = append(newOutput, f(v))
	}
	return newOutput
}