package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/acm-icpc-team

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	s.Scan()
	m, _ := strconv.Atoi(s.Text())

	mb := make([][]bool, n)
	for i := 0; i < n; i++ {
		mb[i] = make([]bool, m)
	}

	s.Split(bufio.ScanLines)

	for i := 0; i < n; i++ {
		s.Scan()
		line := s.Text()
		for j := 0; j < m; j++ {
			switch line[j : j+1] {
			case "1":
				mb[i][j] = true
			case "0":
				mb[i][j] = false
			}
		}
	}

	max := 0
	countMax := 0

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			v := combine(mb[i], mb[j])
			if max < v {
				countMax = 1
				max = v
			} else if max == v {
				countMax++
			}
		}
	}

	fmt.Printf("%v\n", max)
	fmt.Printf("%v\n", countMax)
}

func combine(m1 []bool, m2 []bool) int {
	sum := 0
	for i := 0; i < len(m1); i++ {
		if m1[i] || m2[i] {
			sum++
		}
	}
	return sum
}
