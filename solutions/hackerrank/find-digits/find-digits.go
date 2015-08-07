package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/find-digits

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	for i := 0; i < n; i++ {
		s.Scan()
		x, _ := strconv.Atoi(s.Text())
		cnt := 0
		p := 1
		for p < x {
			v := (x / p) % 10
			if v != 0 && x%v == 0 {
				cnt++
			}
			p *= 10
		}
		fmt.Printf("%v\n", cnt)
	}
}
