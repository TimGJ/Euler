package main

import "fmt"

/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

func Collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}

func CollatzSequence(n int) []int {

	sequence := []int{}
	if n > 0 {
		for {
			sequence = append(sequence, n)
			if n == 1 {
				break
			}
			n = Collatz(n)
		}
	}
	return sequence
}

func main() {
	const Limit = 1000000
	MaxLen := 0
	for n := 0; n < Limit; n++ {
		s := CollatzSequence(n)
		if len(s) > MaxLen {
			MaxLen = len(s)
			fmt.Println(n, MaxLen, s)
		}
	}
}
