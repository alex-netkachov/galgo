package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/service-lane

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	sl := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		sl[i], _ = strconv.Atoi(s.Text())
	}

	for i := 0; i < t; i++ {
		s.Scan()
		in, _ := strconv.Atoi(s.Text())
		s.Scan()
		out, _ := strconv.Atoi(s.Text())
		fmt.Printf("%v\n", getWidth(sl, in, out))
	}
}

func getWidth(sl []int, i int, j int) int {
	min := 100000
	for idx := i; idx <= j; idx += 1 {
		if min > sl[idx] {
			min = sl[idx]
		}
	}
	return min
}
