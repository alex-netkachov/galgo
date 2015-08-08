package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt sherlock-and-watson.go && go run sherlock-and-watson.go < sherlock-and-watson.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	s.Scan()
	q, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}

	a = rot(a, k)

	for i := 0; i < q; i++ {
		s.Scan()
		v, _ := strconv.Atoi(s.Text())
		Println(a[v])
	}
}

func rot(a []int, k int) []int {
	r := make([]int, len(a))
	q := k % len(a)
	for i, j := 0, q; i < len(a); i, j = i+1, j+1 {
		r[j%len(a)] = a[i]
	}
	return r
}
