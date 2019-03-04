package luhn

// IsEven checks whether a number is even
func IsEven(num int) bool {
	return num%2 == 0
}

// GetInt whatever
func GetInt(intval int) int {
	if 48 >= intval && intval <= 57 {
		return intval - 48
	}
	return 0
}

// Valid …ates credit card numbers
func Valid(ccString string) bool {
	if len(ccString) == 1 {
		return false
	}
	sum := 0
	if IsEven(len(ccString)) {
		for p, char := range ccString {
			if !IsEven(p) {
				number := GetInt(int(char)) * 2
				if number > 9 {
					number -= 9
				}
				sum += number
			}
		}
	} else {
		for p, char := range ccString {
			if IsEven(p) {
				number := GetInt(int(char)) * 2
				if number > 9 {
					number = number - 9
				}
				sum = sum + number
			}
		}
	}
	return sum%10 == 0
}
