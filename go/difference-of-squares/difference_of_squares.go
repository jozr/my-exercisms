package diffsquares

// SquareOfSum squares the sum of all natural numbers up to n
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2 // Gauss sequence formula
	return sum * sum
}

// SumOfSquares sums the squares of all natural numbers up to n
func SumOfSquares(n int) int {
	// https://math.stackexchange.com/questions/183316/how-to-get-to-the-formula-for-the-sum-of-squares-of-first-n-numbers
	return (n * (n + 1) * ((2 * n) + 1)) / 6
}

// Difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
