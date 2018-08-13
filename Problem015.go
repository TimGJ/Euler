package main

import (
	"fmt"
	"github.com/TimGJ/Euler/euler"
)

func main() {
	const Size int64 = 13

	steps := Size * 2
	fmt.Println(euler.Combinations(steps, Size))

}
