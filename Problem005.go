package main

import (
	"github.com/TimGJ/Euler/euler"
	"fmt"
	"time"
)

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
//
// We could, of course, burte force this. But I think what we need to do is to get the prime factors of
// all the numbers from 2..20 and then this number will be the product of those prime factors raised to the
// maximum number of times they occur. I think.

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

func Count(f []int) Counter {
	// Returns a count of the incidences of particular values in a slice of ints. So given the input
	// [1, 2, 3, 4, 4, 9, 4, 3, 9, 5] is will return [1:1, 2:1, 3:2, 4:3, 5:1, 9:2]

	m := make(Counter)
	for _, e := range f {
		m[e] += 1
	}
	return m
}

func Consolidate(m MetaCounter) Counter {

	c := make(Counter)

	for key := 0 ; key <= len(m) ; key++ {
		value := m[key]
		for factor, count := range value {
			if c[factor] < count {
				c[factor] = count
			}
		}
	}
	return c
}

func Power(x, n int) int {
	v := 1
	for i := 0 ; i < n ; i++ {
		v *= x
	}
	return v
}

func main() {

	const limit = 20
	m := make(MetaCounter)

	start := time.Now()

	for  v := 1 ; v <= limit ; v++ {
		f := euler.PrimeFactors(v)
		m[v] = Count(f)
	}
	c := Consolidate(m)
	value := 1
	for factor, power := range c {
		value *= Power(factor, power)
	}
	// Check that value is divisible by the numbers 1..limit
	fmt.Println(value)
	for v := 1 ; v <= limit ; v++ {
		if value % v != 0 {
			fmt.Printf("Error! %v  does not divide by %v\n", value, v)
		}
	}
	fmt.Printf("Execution took %v", time.Now().Sub(start))
}

