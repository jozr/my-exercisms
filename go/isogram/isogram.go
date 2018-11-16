package isogram

import "unicode"

// IsIsogram verifies a word has no duplicate letters
func IsIsogram(word string) bool {
	var charCountMap = map[rune]int{}
	boolean := true

	for _, char := range word {
		if unicode.IsLetter(char) {
			charCountMap[unicode.ToLower(char)]++
		}
	}
	for _, count := range charCountMap {
		if count > 1 {
			boolean = false
		}
	}
	return boolean
}
