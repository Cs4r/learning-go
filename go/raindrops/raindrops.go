// Package raindrops converts an integer into "raindrop-speak"
package raindrops

import "strconv"

// ensure that the test version is correct
const testVersion = 3

// Convert does the actual conversion from integer to string
func Convert(num int) string {
	outputString := ""
	if num%3 == 0 {
		outputString += "Pling"
	}
	if num%5 == 0 {
		outputString += "Plang"
	}
	if num%7 == 0 {
		outputString += "Plong"
	}
	if outputString == "" {
		outputString = strconv.Itoa(num)
	}
	return outputString
}