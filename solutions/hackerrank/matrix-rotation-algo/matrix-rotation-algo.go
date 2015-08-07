package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/matrix-rotation-algo

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	m, _ := strconv.Atoi(s.Text())

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Scan()
	r, _ := strconv.Atoi(s.Text())

	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			s.Scan()
			a[i][j], _ = strconv.Atoi(s.Text())
		}
	}

	rotateall(a, m, n, r)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			Printf("%v ", a[i][j])
		}
		Println()
	}
}

func rotateall(a [][]int, m, n, r int) {
	l := m
	if m > n {
		l = n
	}
	l /= 2

	for i := 0; i < l; i++ {
		rotate(a, i, i, m-1-i, n-1-i, r)
	}
}

func rotate(a [][]int, i1, j1, i2, j2, r int) {
	t := save(a, i1, j1, i2, j2)
	shift(t, r)
	load(a, i1, j1, i2, j2, t)
}

func save(a [][]int, i1, j1, i2, j2 int) []int {
	length := (i2 - i1 + j2 - j1) * 2
	tmp := make([]int, length)
	p := 0
	// ----->
	for j := j1; j < j2; j++ {
		tmp[p] = a[i1][j]
		p++
	}
	// |.
	for i := i1; i < i2; i++ {
		tmp[p] = a[i][j2]
		p++
	}
	// <-----
	for j := j2; j > j1; j-- {
		tmp[p] = a[i2][j]
		p++
	}
	// |'
	for i := i2; i > i1; i-- {
		tmp[p] = a[i][j1]
		p++
	}

	return tmp
}

func load(a [][]int, i1, j1, i2, j2 int, tmp []int) {
	p := 0
	// ----->
	for j := j1; j < j2; j++ {
		a[i1][j] = tmp[p]
		p++
	}
	// |.
	for i := i1; i < i2; i++ {
		a[i][j2] = tmp[p]
		p++
	}
	// <-----
	for j := j2; j > j1; j-- {
		a[i2][j] = tmp[p]
		p++
	}
	// |'
	for i := i2; i > i1; i-- {
		a[i][j1] = tmp[p]
		p++
	}
}

func shift(a []int, r int) {
	t := make([]int, len(a))
	offset := r % len(a)
	if offset == 0 {
		copy(t, a)
	} else {
		for i := 0; i < len(a); i++ {
			t[i] = a[(i+offset)%len(t)]
		}
	}
	copy(a, t)
}
