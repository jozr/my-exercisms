package diffsquares

import "math"

// SquareOfSum comment
func SquareOfSum(number int) int {
	sum := 0
	numbers := makeRange(0, number)
	for _, num := range numbers {
		sum += num
	}
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares comment
func SumOfSquares(number int) int {
	var summedSquare int
	ranger := makeRange(0, number)
	for i := range ranger {
		summedSquare += int(math.Pow(float64(i), 2))
	}
	return summedSquare
}

// Difference comment
func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
