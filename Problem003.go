package main

import "fmt"


func Problem003() {
	fmt.Println("Problem #3")
	p := PrimeFactors(600851475143)
	fmt.Println(p[len(p)-1])
}
