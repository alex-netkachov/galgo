package main

import (
	. "fmt"
	"sort"
)

// go fmt closest-numbers.go && go run closest-numbers.go < closest-numbers.txt

func main() {
	n := 0
	Scanln(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		Scan(&a[i])
	}

	sort.Ints(a)

	min := abs(a[n-1] - a[0])
	for i := 1; i < n; i++ {
		d := abs(a[i-1] - a[i])
		if d < min {
			min = d
		}
	}

	for i := 1; i < n; i++ {
		d := abs(a[i-1] - a[i])
		if d == min {
			Printf("%v %v ", a[i-1], a[i])
		}
	}
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
