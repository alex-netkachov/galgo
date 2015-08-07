package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/gem-stones

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanLines)

	es := make(map[rune]int)
	for i := 0; i < n; i++ {
		s.Scan()
		rock := s.Text()
		les := make(map[rune]bool)
		for _, i := range rock {
			les[i] = true
		}
		for e, _ := range les {
			v, _ := es[e]
			es[e] = v + 1
		}
	}

	cnt := 0
	for _, c := range es {
		if c == n {
			cnt++
		}
	}

	fmt.Printf("%v\n", cnt)
}
