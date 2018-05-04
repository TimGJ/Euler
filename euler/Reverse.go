package euler

import (
	"strconv"
	"fmt"
)

func ReverseString(s string) string {
	// Reverses a string

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseNumber(n int) int {
	// Reverses a base 10 integer. So if n is 1234, the function will return 4321

	s := strconv.Itoa(n)

	if r, e := strconv.Atoi(ReverseString(s)); e == nil {
		return r
	} else {
		fmt.Println(e.Error())
		return 0
	}

}

