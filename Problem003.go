package main

import (
	"fmt"
	"github.com/TimGJ/Euler/euler"
)


func main() {
	fmt.Println("Problem #3")
	p := euler.PrimeFactors(600851475143)
	fmt.Println(p[len(p)-1])
}

