package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/cavity-map

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanRunes)

	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s.Scan()
			v, err := strconv.Atoi(s.Text())
			for err != nil {
				s.Scan()
				v, err = strconv.Atoi(s.Text())
			}
			m[i][j] = v
		}
	}

	x := make([]int, 0)
	y := make([]int, 0)

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			v := m[i][j]
			if m[i-1][j] < v && m[i+1][j] < v && m[i][j-1] < v && m[i][j+1] < v {
				x = append(x, i)
				y = append(y, j)
			}
		}
	}

	ms := make([][]string, n)
	for i := 0; i < n; i++ {
		ms[i] = make([]string, n)
		for j := 0; j < n; j++ {
			ms[i][j] = fmt.Sprintf("%v", m[i][j])
		}
	}

	for i := 0; i < len(x); i++ {
		ms[x[i]][y[i]] = "X"
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v", ms[i][j])
		}
		fmt.Printf("\n")
	}
}
