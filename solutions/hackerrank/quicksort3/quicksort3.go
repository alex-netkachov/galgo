package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt quicksort3.go && go run quicksort3.go < quicksort3.txt

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

	quicksort(a, 0, len(a)-1)
}

func quicksort(a []int, lo, hi int) {
	if lo < hi {
		p := partition(a, lo, hi)
		quicksort(a, lo, p-1)
		quicksort(a, p+1, hi)
	}
}

func partition(a []int, lo, hi int) int {
	pivotIndex := choosePivot(a, lo, hi)
	pivotValue := a[pivotIndex]
	a[pivotIndex], a[hi] = a[hi], a[pivotIndex]
	storeIndex := lo

	for i := lo; i <= hi-1; i++ {
		if a[i] < pivotValue {
			a[i], a[storeIndex] = a[storeIndex], a[i]
			storeIndex++
		}
	}

	a[storeIndex], a[hi] = a[hi], a[storeIndex]

	print(a)

	return storeIndex
}

func choosePivot(a []int, lo, hi int) int {
	return hi
}

func sprint(a []int) string {
	str := ""
	for i, v := range a {
		str += fmt.Sprintf("%v", v)
		if i != len(a)-1 {
			str += " "
		}
	}
	return str
}

func print(a []int) {
	str := sprint(a)
	fmt.Printf("%v\n", str)
}
