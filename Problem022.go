package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

/*
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

func ScoreString(s string) int {
	var total int
	runes := []rune(s)
	for _, token := range runes {
		value := int(token-'A') + 1
		total += value
	}
	return total
}

func main() {
	const filename = "p022_names.txt"
	var total int
	if buffer, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		lines := strings.Split(string(buffer), ",")
		for i := 0; i < len(lines); i++ {
			lines[i] = strings.Replace(lines[i], "\"", "", -1)
		}
		sort.Strings(lines)
		for i, s := range lines {
			total += (i + 1) * ScoreString(s)
		}
		fmt.Println(total)
	}
}
