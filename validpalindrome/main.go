package main

import (
	"fmt"
	"strings"
)

var alphaMap map[byte]string

func main() {
	//atozs := []byte("abcdefghijklmnopqrstuvwxyz")
	//AtoZs := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// fmt.Println(isPalindrome("aa"))
	// fmt.Println(isPalindrome("aaa"))
	// fmt.Println(isPalindrome("aabb"))
	// fmt.Println(isPalindrome("abba"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	// fmt.Println(reverse("abba"))
	// fmt.Println(reverse("abcd"))
}
func isPalindrome(s string) bool {
	chars := []rune(s)
	len := len(s)
	left := 0
	right := len - 1

	for left < right {
		for isAlphaNumeric(chars[left]) == false {
			left++
			if left == right {
				return true
			}
		}

		for isAlphaNumeric(chars[right]) == false {
			right--
			if left == right {
				return true
			}
		}

		if toLowerCase(chars[left]) != toLowerCase(chars[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphaNumeric(c rune) bool {
	if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) || (c >= 48 && c <= 57) {
		return true
	}
	return false
}

func toLowerCase(c rune) rune {
	if c >= 65 && c <= 90 {
		return c + 32
	}
	return c
}

func isPalindromeSlow(s string) bool {

	var replacer = strings.NewReplacer(
		" ", "",
		".", "",
		":", "",
		",", "",
		"@", "",
		"#", "",
		"-", "",
	)
	s = strings.ToLower(replacer.Replace(s))

	if len(s) == 0 {
		return true
	}

	if len(s) == 1 {
		return true
	}

	if len(s) == 2 {
		return s[0] == s[1]
	}

	return s == reverse(s)
}

func reverse(s string) string {
	var tmp rune
	strArr := []rune(s)
	slen := len(s)
	head := 0
	tail := slen - 1

	for head != tail {
		tmp = strArr[head]
		strArr[head] = strArr[tail]
		strArr[tail] = tmp
		if head+1 == tail {
			break
		} else {
			head++
			tail--
		}
	}
	return string(strArr)
}
