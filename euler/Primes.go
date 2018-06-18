package euler

import "fmt"

// Various functions which are likely to be useful in the solution of problems in Project Euler
//
// Doing ithis more to get my head around how Go packages work

// Get all prime factors of a given number n
func PrimeFactors(n int) []int {
	// Get the number of 2s that divide n
	pfs := []int{}
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}
	return pfs
}

// Counters for counting up prime factors.

type Counter map[int]int

func (c Counter) String() string {

	s := ""

	for key, value := range c {
		s += fmt.Sprintf("%d: (%d)\t", key, value)
	}
	return s
}

type MetaCounter map[int]Counter

func (c MetaCounter) String() string {

	s := ""

	for key, value := range c {
		s += fmt.Sprintf("%d: [%v]\t", key, value)
	}
	return s
}

func GetFirstNPrimes(n int) []int {
	primes := []int{}

	primes = append(primes, 2)

	for m := 3; len(primes) < n; m++ {
		isprime := true
		for _, p := range primes {
			if m%p == 0 {
				isprime = false
			}
		}
		if isprime {
			primes = append(primes, m)
		}
	}
	return primes
}

func GetPrimesUpto(n int) []int {
	// Gets all prime numbers less than n
	primes := []int{}

	if n < 2 {
		return primes
	}

	primes = append(primes, 2)

	for m := 3; m <= n; m++ {
		isprime := true
		for _, p := range primes {
			if m%p == 0 {
				isprime = false
			}
		}
		if isprime {
			primes = append(primes, m)
		}
	}
	return primes
}
