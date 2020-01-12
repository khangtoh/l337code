package main

import "fmt"

func main() {

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))

	fmt.Println(lengthOfLongestSubstring("aaaa"))
	fmt.Println(lengthOfLongestSubstring("abcc"))

	fmt.Println(lengthOfLongestSubstring("pwwkesw"))

}

func lengthOfLongestSubstring(s string) int {

	// visited contains the index of the char when it was visited
	// init to -1
	//length := len(s)
	m := make(map[rune]int)
	start := 0
	currentLen := 1
	maxLen := 1

	// first char is always unique
	for _, char := range s {
		m[char] = -1
	}
	for pos, char := range s {

		if m[char] == -1 {
			m[char] = pos
			if len(s)-pos == 1 && pos-start > currentLen {
				currentLen = pos - start
			}
		} else {

			if len(s)-pos > 1 {
				if pos-start > currentLen {
					currentLen = pos - start
				}
				m = make(map[rune]int)
				start = pos
			} else {
				currentLen = pos - start
			}
		}
	}
	if currentLen > maxLen {
		maxLen = currentLen
	}
	return maxLen
}
