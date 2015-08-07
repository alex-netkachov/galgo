package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/chocolate-feast

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	for i := 0; i < n; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		s.Scan()
		c, _ := strconv.Atoi(s.Text())
		s.Scan()
		m, _ := strconv.Atoi(s.Text())

		cnt := n / c
		w := n / c
		for w >= m {
			cnt += w / m
			w = w/m + w%m
		}

		fmt.Printf("%v\n", cnt)
	}
}
