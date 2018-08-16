package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Maximum path sum I
Problem 18
By starting at the top of the triangle below and moving to adjacent numbers on the Row below, the maximum total from top to bottom is 23.

   3
  7 4
 2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle below:

                            75
                          95  64
                        17  47  82
                      18  35  87  10
                    20  04  82  47  65
                  19  01  23  75  03  34
                88  02  77  73  07  63  67
              99  65  04  28  06  16  70  92
            41  41  26  56  83  40  80  70  33
          41  48  72  33  47  32  37  16  94  29
        53  71  44  65  25  43  91  52  97  51  14
      70  11  33  28  77  73  17  78  39  68  17  57
    91  71  52  38  17  14  91  43  58  50  27  29  48
  63  66  04  68  89  53  67  30  73  16  69  87  40  31
04  62  98  27  23  09  70  98  73  93  38  53  60  04  23

NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route.
However, Problem 67, is the same challenge with a triangle containing one-hundred rows;
it cannot be solved by brute force, and requires a clever method! ;o)
*/

type Node struct {
	value               int64
	left, right, parent *Node
	chosen              bool
}

func (n Node) String() string {
	bob := strings.Builder{}
	bob.WriteString(fmt.Sprintf("%d:\t", n.value))
	if n.left == nil {
		bob.WriteString("Left -> nil\t")
	} else {
		bob.WriteString(fmt.Sprintf("Left -> %d\t", n.left.value))
	}

	if n.right == nil {
		bob.WriteString("Right -> nil\t")
	} else {
		bob.WriteString(fmt.Sprintf("Right -> %d\t", n.right.value))
	}

	if n.parent == nil {
		bob.WriteString("Parent -> nil\t")
	} else {
		bob.WriteString(fmt.Sprintf("Parent -> %d\t", n.parent.value))
	}
	return bob.String()
}

type Row []Node

func (r Row) String() string {
	s := []string{}
	for _, a := range r {
		s = append(s, fmt.Sprintf("%d", a.value))
	}

	return fmt.Sprintf("%d elements: (%s)", len(s), strings.Join(s, ", "))
}

type Matrix []Row

func (m Matrix) String() string {
	bob := strings.Builder{}

	rows := len(m)
	for row := 0; row < rows; row++ {
		spaces := 2 * (rows - row)
		bob.WriteString(strings.Repeat(" ", spaces))
		numbers := []string{}
		for _, number := range m[row] {
			numbers = append(numbers, fmt.Sprintf("%2d", number.value))
		}
		bob.WriteString(strings.Join(numbers, "  "))
		bob.WriteString("\n")
	}
	return bob.String()
}

func (m Matrix) MakeLinks() {
	for row := 0; row < len(m)-1; row++ {
		for cell := 0; cell < len(m[row]); cell++ {
			m[row][cell].left = &m[row+1][cell]
			m[row][cell].right = &m[row+1][cell+1]
			m[row][cell].left.parent = &m[row][cell]
			m[row][cell].right.parent = &m[row][cell]
			fmt.Printf("(%d, %d): %v \n", row, cell, m[row][cell])
		}
	}
}

func MakeRow(a []string) Row {
	r := Row{}
	for _, s := range a {
		var i int64
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		r = append(r, Node{value: i})
	}
	return r
}

func main() {

	input := `
   3
  7 4
 2 4 6
8 5 9 3
`

	m := Matrix{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		row := MakeRow(fields)
		m = append(m, row)
	}
	m.MakeLinks()
}
