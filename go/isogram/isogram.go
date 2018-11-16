package isogram

import "unicode"

// IsIsogram verifies a word has no duplicate letters
func IsIsogram(word string) bool {
	var charCountMap = map[rune]int{}
	for _, char := range word {
		lowercaseChar := unicode.ToLower(char)
		if unicode.IsLetter(char) {
			charCountMap[lowercaseChar]++
		}
		if charCountMap[lowercaseChar] > 1 {
			return false
		}
	}
	return true
}
