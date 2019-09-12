//Package diffsquares determines the square of the sum of n positive integers,
//the sum of the square of each n positive integer,
//and the difference between both.
package diffsquares

//SquareOfSum = (1+2+3+...+n)^2 = n^3 + (n-1)^3 + (n-2)^3 + ... + (n-n)^3.
//SquareOfSum = n^3 + {n*(n-1)/2}^2
func SquareOfSum(n int) int {

	return n*n*n + (n*(n-1)/2)*(n*(n-1)/2)

}

//SumOfSquares = (1^2 + 2^2 + ... + n^2) = n^2 + (n-1)^2 + (n-2)^2 + ... + (n-n)^2
// (n/6)*(n+1)*(2n+1)
func SumOfSquares(n int) int {

	return n * (n + 1) * (2*n + 1) / 6
}

//Difference = SquareOfSum - SumOfSquares.
func Difference(n int) int {

	return SquareOfSum(n) - SumOfSquares(n)
}
