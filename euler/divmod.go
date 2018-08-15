package euler

import (
	"log"
)

func DivMod(a, b int) (int, int) {
	if b == 0 {
		log.Fatal("Divide by zero!")
	}
	return a / b, a % b
}
