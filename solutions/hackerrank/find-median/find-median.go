package main

import (
	. "fmt"
)

// go fmt find-median.go && go run find-median.go < find-median.txt

func main() {
	n := 0
	Scanln(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		Scan(&a[i])
	}

	median(a, 0, len(a)-1)
	Printf("%v\n", a[len(a)/2])
}

func median(a []int, i, j int) {
	if i < j {
		p := partition(a, i, j)
		mi := len(a) / 2
		if p > mi {
			median(a, i, p-1)
		}
		if p < mi {
			median(a, p+1, j)
		}
	}
}

func partition(a []int, i, j int) int {
	p := pivot(a, i, j)
	pv := a[p]
	a[j], a[p] = a[p], a[j]
	m := i
	for k := i; k < j; k++ {
		if a[k] < pv {
			a[k], a[m] = a[m], a[k]
			m++
		}
	}
	a[m], a[j] = a[j], a[m]
	return m
}

func pivot(a []int, i, j int) int {
	return i
}
