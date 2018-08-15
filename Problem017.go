package main

import (
	"fmt"
	"github.com/TimGJ/Euler/euler"
	"log"
	"strings"
	"unicode"
)

/*
** If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
**
** If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
 */

func NumberToText(n int) string {
	/*
	** Converts a number into text format
	 */
	bob := strings.Builder{}
	var andreq bool
	names := map[int]string{0: "zero", 1: "one", 2: "two", 3: "three", 4: "four",
		5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
		10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
		15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen",
		20: "twenty", 30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety",
		100: "hundred", 1000: "thousand"}

	if n >= 10000 {
		log.Fatal("Number must be < 10000")
	}

	thousands, remainder := euler.DivMod(n, 1000)
	if thousands > 0 {
		bob.WriteString(fmt.Sprintf("%s %s ", names[thousands], names[1000]))
		andreq = true
	}
	hundreds, remainder := euler.DivMod(remainder, 100)
	if hundreds > 0 {
		bob.WriteString(fmt.Sprintf("%s %s ", names[hundreds], names[100]))
		andreq = true
	}

	if remainder != 0 && andreq {
		bob.WriteString("and ")
	}
	tens, remainder := euler.DivMod(remainder, 10)
	if tens > 1 {
		bob.WriteString(fmt.Sprintf("%s ", names[tens*10]))
	} else if tens == 1 {
		remainder += tens * 10
	}
	if remainder != 0 {
		bob.WriteString(fmt.Sprintf("%s ", names[remainder]))
	}
	return bob.String()
}

func StripSpace(s string) string {
	t := []rune(s)
	u := []rune{}
	for _, v := range t {
		if !unicode.IsSpace(v) {
			u = append(u, v)
		}
	}
	return string(u)
}

func main() {
	bob := strings.Builder{}
	for value := 1; value <= 1000; value++ {
		bob.WriteString(NumberToText(value))
	}
	s := StripSpace(bob.String())
	fmt.Printf("%d: <%s>\n", len(s), s)

}
