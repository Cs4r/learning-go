// Package hamming calculates the Hamming difference between two DNA strands.
package hamming

import "errors"

// ensure correct test version
const testVersion = 6

// Distance iterates through each sequence in order and tallies up any differences.
func Distance(a, b string) (int, error) {
	numDifferences := 0
	if len(a) != len(b) {
		return -1, errors.New("DNA strands must be of equal length.")
	}
	i := 0
	for ; i < len(a); i++ {
		if a[i] != b[i] {
			numDifferences++
		}
	}
	return numDifferences, nil
}