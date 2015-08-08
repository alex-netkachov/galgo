package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt quicksort2.go && go run quicksort2.go < quicksort2.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}

	print(partition(a))
}

func partition(a []int) []int {
	if len(a) < 2 {
		return a
	}
	if len(a) == 2 {
		if a[0] < a[1] {
			return a
		}
		return []int{a[1], a[0]}
	}

	p := a[0]
	al := []int{}
	ar := []int{}
	for i := 1; i < len(a); i++ {
		v := a[i]
		if v > p {
			ar = append(ar, v)
		} else {
			al = append(al, v)
		}
	}
	if len(al) > 1 {
		al = partition(al)
		print(al)
	}
	if len(ar) > 1 {
		ar = partition(ar)
		print(ar)
	}
	return append(append(al, p), ar...)
}

func print(a []int) {
	for _, v := range a {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
}
