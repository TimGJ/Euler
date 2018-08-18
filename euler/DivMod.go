package euler

import (
	"log"
)

func DivMod(a, b int) (int, int) {
	if b == 0 {
		log.Fatal("Divide by zero!")
	}
	return a / b, a % b
}

func ProperDivisors(n int) []int {

	/*
		Returns all proper divisors for n. So for n == 22, would return [1, 2, 11]
	*/
	divisors := []int{1}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func SumProperDivisors(n int) int {
	p := ProperDivisors(n)
	var sum int

	for _, k := range p {
		sum += k
	}
	return sum
}
