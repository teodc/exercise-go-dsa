package other

import (
	"regexp"
	"strings"
)

func IsPalindrome(word string) bool {
	// lower the case
	cleanWord := strings.ToLower(word)
	// remove spaces
	cleanWord = strings.Join(strings.Fields(cleanWord), "")
	// remove non alpha-num characters
	var nonAlphanumRegex = regexp.MustCompile("[^a-zA-Z0-9 ]+")
	cleanWord = nonAlphanumRegex.ReplaceAllString(cleanWord, "")

	var i, j int
	chars := []rune(cleanWord)
	for i = 0; i < len(chars)/2; i++ {
		j = len(chars) - 1 - i
		if string(chars[i]) != string(chars[j]) {
			return false
		}
	}

	return true
}
