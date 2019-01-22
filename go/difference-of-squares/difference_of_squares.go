package diffsquares

import "math"

// SquareOfSum squares the sum of all numbers up to n
func SquareOfSum(number int) int {
	sum := 0
	numbers := rangeFromZero(number)
	for num := range numbers {
		sum += num
	}
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares sums the squares of all numbers up to n
func SumOfSquares(number int) int {
	var summedSquare int
	numbers := rangeFromZero(number)
	for num := range numbers {
		summedSquare += int(math.Pow(float64(num), 2))
	}
	return summedSquare
}

// Difference between SquareOfSum and SumOfSquares
func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}

func rangeFromZero(max int) []int {
	rangeArray := make([]int, max+1)
	for num := range rangeArray {
		rangeArray[num] = num
	}
	return rangeArray
}
