package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt largest-permutation.go && go run largest-permutation.go < largest-permutation.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}

	for i, j := 0, 0; i < n-1 && j < k; i++ {
		idxs := maxidxs(a, i)
		if -1 == indexOf(idxs, i) {
			f := idxs[len(idxs)-1]
			a[f], a[i] = a[i], a[f]
			j++
		}
	}

	for i := 0; i < n; i++ {
		Printf("%v ", a[i])
	}
	Println()
}

func rev(a []int) []int {
	r := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		r[i] = a[len(a)-1-i]
	}
	return r
}

func maxidxs(a []int, s int) []int {
	b := a[s:]
	m := b[0]
	for _, v := range b {
		if m < v {
			m = v
		}
	}

	r := make([]int, 0)
	for i, v := range b {
		if v == m {
			r = append(r, i+s)
		}
	}
	return r
}

func indexOf(a []int, n int) int {
	for i, v := range a {
		if v == n {
			return i
		}
	}
	return -1
}
