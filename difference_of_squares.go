//Package diffsquares determines the square of the sum of n positive integers,
//the sum of the square of each n positive integer, and the difference between both.
package diffsquares

//SquareOfSum = (1+2+3+...+n)^2 = n^3 + (n-1)^3 + (n-2)^3 + ... + (n-n)^3.
func SquareOfSum(n int) int {
	if n == 1 {
		return n * n * n
	}

	return n*n*n + SquareOfSum(n-1)

}

//SumOfSquares = (1^2 + 2^2 + ... + n^2) = n^2 + (n-1)^2 + (n-2)^2 + ... + (n-n)^2.
func SumOfSquares(n int) int {

	if n == 1 {
		return n * n
	}
	return n*n + SumOfSquares(n-1)
}

//Difference = SquareOfSum - SumOfSquares = n^2(n-1)+(n-1)^2(n-2)+...+1^2(n-n).
func Difference(n int) int {

	if n == 1 {
		return n * n * (n - 1)
	}
	return n*n*(n-1) + Difference(n-1)
}
