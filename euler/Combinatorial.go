package euler

/*
Various combinatorial functions
*/

func Factorial(n int64) int64 {
	var factorial int64 = 1
	var i int64

	for i = 1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}

func Combinations(n, r int64) int64 {
	if r > n {
		n, r = r, n
	}

	var t int64 = 1
	for v := n; v > r; v-- {
		t *= v
	}
	b := Factorial(n - r)
	return t / b
}
