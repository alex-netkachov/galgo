package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/utopian-tree

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	for i := 0; i < n; i++ {
		s.Scan()
		t, _ := strconv.Atoi(s.Text())
		fmt.Printf("%v\n", grow(t))
	}
}

func grow(n int) int {
	height := 1
	i := 0
	for i < n {
		height *= 2
		i++
		if i == n {
			break
		}
		height++
		i++
		if i == n {
			break
		}
	}
	return height
}
