package main

import (
	"fmt"
	"strings"
)

/*

Return the lexicographically smallest subsequence of text that contains all the distinct characters of text exactly once.

Example 1:

Input: "cdadabcc"
Output: "adbc"

Example 2:

Input: "abcd"
Output: "abcd"

Example 3:

Input: "ecbacba"
Output: "eacb"

Example 4:

Input: "leetcode"
Output: "letcod"

*/

func main() {
	fmt.Print("main")
	fmt.Print(smallestSubsequence("cdadabcc"))
}

func smallestSubsequence(text string) string {
	distinct := 0            // count the number of distinct letters
	count := make([]int, 26) // counting frequency of each letters
	for i := 0; i < len(text); i++ {
		count[text[i]-'a']++
		if count[text[i]-'a'] == 1 {
			distinct++
		}
	}

	res := make([]rune, distinct)
	res[0] = rune(text[0]) // push the first letter into the stack
	count[res[0]-'a']--
	p := 0 // pointer of the stack
	for i := 1; i < len(text); i++ {
		count[text[i]-'a']--
		if strings.Contains(string(res[0:p+1]), string(text[i])) {
			continue
		} else {
			for p >= 0 && rune(text[i]) < res[p] && count[res[p]-'a'] > 0 {
				p--
			}
			p++
			res[p] = rune(text[i])
		}
	}
	return string(res)
}
