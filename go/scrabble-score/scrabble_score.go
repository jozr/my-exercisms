package scrabble

import "strings"

// Score words, scrabble style
func Score(word string) int {
	var scores = map[int][]string{
		1:  []string{"a", "e", "i", "o", "u", "l", "n", "r", "s", "t"},
		2:  []string{"d", "g"},
		3:  []string{"b", "c", "m", "p"},
		4:  []string{"f", "h", "v", "w", "y"},
		5:  []string{"k"},
		8:  []string{"j", "x"},
		10: []string{"q", "z"},
	}
	finalScore := 0
	for _, letter := range word {
		for score := range scores {
			for _, char := range scores[score] {
				if strings.ToLower(string(letter)) == char {
					finalScore = finalScore + score
				}
			}
		}
	}
	return finalScore
}
