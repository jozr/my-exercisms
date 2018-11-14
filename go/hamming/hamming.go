package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(b) != len(a) {
		return -1, errors.New("strings 'a' and 'b' are unequal in length")
	}
	distance := 0
	for i, _ := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
