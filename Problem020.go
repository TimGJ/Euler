package main

import (
	"fmt"
	"math/big"
)

/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

func DigitSum(k *big.Int) *big.Int {
	// Sums the digits
	total := big.NewInt(0)
	mod := big.NewInt(0)
	ten := big.NewInt(10)
	for k.Sign() > 0 {
		k.DivMod(k, ten, mod)
		total.Add(total, mod)
		fmt.Println(k, mod, total)
	}
	return total
}
func main() {

	n := big.NewInt(100)
	p := big.NewInt(1)
	f := big.NewInt(1)
	for i := big.NewInt(2); i.Cmp(n) < 1; i.Add(i, p) {
		f.Mul(f, i)
		fmt.Println(i, f)
	}
	fmt.Println(DigitSum(f))
}
