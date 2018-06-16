package main

import (
	"fmt"
	"math"
)

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a**2 + b**2 = c**2
For example, 3**2 + 4**2 = 9 + 16 = 25 = 5**2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

func main() {

	const sum = 1000

	for a := 1; a < sum; a++ {
		a2 := a * a
		for b := a; b < sum; b++ {
			b2 := b * b
			c2 := a2 + b2
			c := int(math.Sqrt(float64(c2)))
			if c*c == c2 && (a+b+c) == sum {
				fmt.Println(a, b, c, a*b*c)
			}
		}
	}
}
