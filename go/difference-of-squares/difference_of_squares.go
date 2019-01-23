package diffsquares

// SquareOfSum squares the sum of all numbers up to n
func SquareOfSum(number int) int {
	sum := 0
	for i := 0; i <= number; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares sums the squares of all numbers up to n
func SumOfSquares(number int) int {
	var summedSquare int
	for i := 0; i <= number; i++ {
		summedSquare += i * i
	}
	return summedSquare
}

// Difference between SquareOfSum and SumOfSquares
func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}
