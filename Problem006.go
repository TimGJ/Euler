package main

import "fmt"

/*
The sum of the squares of the first ten natural numbers is,

1**2 + 2**2 + ... + 10**2 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)**2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

func SquareSum(n int) int {
	// Return the square of the sum of the first n natural numbers
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumSquares(n int) int {
	// Return the sum of the squares of the first n natural numbers
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func main() {

	for _, n := range []int{10, 100} {

		a := SquareSum(n)
		b := SumSquares(n)

		fmt.Printf("%d: %d - %d = %d\n", n, a, b, a-b)
	}

}
