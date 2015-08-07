package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://www.hackerrank.com/challenges/common-child

var c [][]int

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	a := s.Text()

	s.Scan()
	b := s.Text()

	c = make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = make([]int, len(b))
		for j := 0; j < len(b); j++ {
			c[i][j] = -1
		}
	}

	max := find(a, 0, b, 0, 0)

	Printf("%v\n", max)
}

func find(a string, i int, b string, j int, m int) int {
	if i == len(a) || j == len(b) {
		return 0
	}

	v := c[i][j]
	if v > -1 {
		return v
	}

	if a[i] == b[j] {
		v = find(a, i+1, b, j+1, m+1) + 1
	} else {
		v1 := find(a, i+1, b, j, m)
		v2 := find(a, i, b, j+1, m)

		if v1 > v2 {
			v = v1
		} else {
			v = v2
		}
	}

	c[i][j] = v
	return v
}
