package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/the-grid-search

func readMatrix(s *bufio.Scanner) [][]byte {
	s.Split(bufio.ScanWords)

	s.Scan()
	r, _ := strconv.Atoi(s.Text())

	s.Scan()
	c, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanLines)

	g := make([][]byte, r)
	for i := 0; i < r; i++ {
		s.Scan()
		t := s.Text()
		g[i] = make([]byte, c)
		for j := 0; j < c; j++ {
			g[i][j] = t[j] - '0'
		}
	}
	return g
}

func find(line []byte, index int, pattern []byte) int {
	for i := index; i < len(line)-len(pattern)+1; i++ {
		if bytes.Equal(pattern, line[i:i+len(pattern)]) {
			return i
		}
	}
	return -1
}

func doesContain(grid [][]byte, pattern [][]byte) bool {
	for i := 0; i < len(grid)-len(pattern)+1; i++ {
		idx := -1
		for {
			idx = find(grid[i], idx+1, pattern[0])
			if idx == -1 {
				break
			}

			ok := true
			for j := 1; j < len(pattern) && ok; j++ {
				ok = bytes.Equal(pattern[j], grid[i+j][idx:idx+len(pattern[j])])
			}
			if ok {
				return true
			}
		}
	}
	return false
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	for t := 0; t < n; t++ {
		g := readMatrix(s)
		p := readMatrix(s)

		if doesContain(g, p) {
			Println("YES")
		} else {
			Println("NO")
		}
	}
}
