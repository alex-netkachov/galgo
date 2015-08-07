package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/plus-minus

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	list := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		list[i], _ = strconv.Atoi(s.Text())
	}

	cntPos := 0.0
	cntNeg := 0.0
	cntZer := 0.0

	for i := 0; i < n; i++ {
		switch {
		case list[i] > 0:
			cntPos++
		case list[i] < 0:
			cntNeg++
		default:
			cntZer++
		}
	}

	fmt.Printf("%.3f\n", cntPos/float64(n))
	fmt.Printf("%.3f\n", cntNeg/float64(n))
	fmt.Printf("%.3f\n", cntZer/float64(n))
}
