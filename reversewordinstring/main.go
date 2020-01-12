package main

import (
	"fmt"
	"strings"
)

// submission on leetcode
// https://leetcode.com/submissions/detail/293528291/

func main() {
	testStr := "the sky is blue"
	fmt.Printf("'%s'", reverseWords(testStr))

	testStr = " Hello! World! "
	fmt.Printf("'%s'", reverseWords(testStr))

	testStr = " Hello!    World! "
	fmt.Printf("'%s'", reverseWords(testStr))
}

func reverseWords(s string) string {

	s = strings.Trim(s, " ")

	var reversed string
	var start, end, slen int
	slen = len(s)
	start = slen
	end = slen - 1
	var foundWord = false

	for i := end; i > 0; i-- {
		if string(s[i]) == " " {

			if foundWord == false {
				if i != 0 && string(s[i-1]) != " " {
					foundWord = false
					start = i + 1
					reversed = reversed + s[start:end+1] + " "
					end = i - 1
				} else {
					foundWord = true
					start = i + 1
				}

			} else if i != 0 && string(s[i-1]) != " " {
				foundWord = false
				reversed = reversed + s[start:end+1] + " "
				end = i - 1
			}

		}
	}
	// copy the last word
	reversed = reversed + s[0:end+1]
	return reversed
}
