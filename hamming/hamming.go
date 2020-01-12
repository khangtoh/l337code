// Package hamming provides a function Distance for determining the hamming distance
// for between 2 strings
package hamming

import (
	"errors"
)

// Distance is a function that determines the hamming distance between 2 strings
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("a is not same length as b")
	}

	var distance int
	for pos := range a {
		if a[pos] != b[pos] {
			distance++
		}
	}

	return distance, nil
}
