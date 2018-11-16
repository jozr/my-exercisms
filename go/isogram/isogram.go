package isogram

import "unicode"

// IsIsogram foo bar baz
func IsIsogram(word string) bool {
	var charCountMap = map[rune]int{}
	boolean := true

	for _, char := range word {
		charCountMap[unicode.ToLower(char)]++
	}
	for char, count := range charCountMap {
		if count > 1 && char != '-' && char != ' ' {
			boolean = false
		}
	}
	return boolean
}
