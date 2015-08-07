package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/cut-the-sticks

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}

	for !done(a) {
		fmt.Printf("%v\n", cut(a))
	}
}

func cut(a []int) int {
	min := 10000
	for _, i := range a {
		if i != 0 && min > i {
			min = i
		}
	}

	cnt := 0
	for i, v := range a {
		if v != 0 {
			v -= min
			if v < 0 {
				v = 0
			}
			a[i] = v
			cnt++
		}
	}

	return cnt
}

func done(a []int) bool {
	for _, i := range a {
		if i != 0 {
			return false
		}
	}
	return true
}
