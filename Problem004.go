package main

import (
	"math"
	"github.com/TimGJ/Euler/euler"
)

//A palindromic number reads the same both ways. The largest palindrome made from the product of two
// 2-digit numbers is 9009 = 91 × 99.
//
//Find the largest palindrome made from the product of two 3-digit numbers.


func main() {

	const digits = 3
	var lower, upper = int(math.Pow10(digits-1)), int(math.Pow10(digits))
	var largest int

	for i := lower ; i < upper ; i++ {
		for j := lower ;  j <= i ; j++ {
			product := i*j
			reversed := euler.ReverseNumber(product)
			if product == reversed && product > largest {
				largest = product
			}
		}
	}
	print("Largest palindromic product of two ", digits, " digit numbers is ", largest)
}
