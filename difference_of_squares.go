//Package diffsquares determines the difference between the sum of squares and the square of sums up to the number n.
package diffsquares

//SquareOfSum = (1+...+n)^2
func SquareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum = sum + i
	}

	return sum * sum
}

//SumOfSquares = 1^2 + ... +n^2
func SumOfSquares(n int) int {
	var sumSqs int
	for i := 1; i <= n; i++ {
		sumSqs = sumSqs + i*i
	}
	return sumSqs
}

//Difference = (1+...+n)^2 - 1^2 - ... - n^2
func Difference(n int) int {
	Difference := SquareOfSum(n) - SumOfSquares(n)
	return Difference
}
