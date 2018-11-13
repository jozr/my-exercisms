package hamming

import "errors"

func Distance(a, b string) (distance int, err error) {
	if len(b) != len(a) {
		return -1, errors.New("Strings 'a' and 'b' are unequal in length.")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
