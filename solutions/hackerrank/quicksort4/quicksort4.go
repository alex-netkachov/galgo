package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt quicksort4.go && go run quicksort4.go < quicksort4.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a1 := make([]int, n)
	a2 := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		v, _ := strconv.Atoi(s.Text())
		a1[i] = v
		a2[i] = v
	}

	fmt.Printf("%v\n", insertsort(a1)-quicksort(a2, 0, len(a2)-1))
}

func insertsort(a []int) int {
	shifts := 0
	j := 0

	for i := 1; i < len(a); i++ {
		value := a[i]
		j = i - 1
		for j > -1 && value < a[j] {
			shifts++
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = value
	}
	return shifts
}

func quicksort(a []int, lo, hi int) int {
	sum := 0
	if lo < hi {
		p, sw := partition(a, lo, hi)
		sum += sw
		sum += quicksort(a, lo, p-1)
		sum += quicksort(a, p+1, hi)
	}
	return sum
}

func partition(a []int, lo, hi int) (int, int) {
	swaps := 0
	pivotIndex := choosePivot(a, lo, hi)
	pivotValue := a[pivotIndex]
	a[pivotIndex], a[hi] = a[hi], a[pivotIndex]
	storeIndex := lo

	for i := lo; i <= hi-1; i++ {
		if a[i] < pivotValue {
			a[i], a[storeIndex] = a[storeIndex], a[i]
			swaps++
			storeIndex++
		}
	}

	a[storeIndex], a[hi] = a[hi], a[storeIndex]
	swaps++

	return storeIndex, swaps
}

func choosePivot(a []int, lo, hi int) int {
	return hi
}
