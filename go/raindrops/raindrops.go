package raindrops

import "fmt"

func Convert(number int) string {
	sound := ""
	if number%3 == 0 {
		sound += "Pling"
	}
	if number%5 == 0 {
		sound += "Plang"
	}
	if number%7 == 0 {
		sound += "Plong"
	}
	if sound == "" {
		return fmt.Sprintf("%d", number)
	}
	return sound
}
