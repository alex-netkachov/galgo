package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt almost-sorted.go && go run almost-sorted.go < almost-sorted.txt

func main() {
	//test()

	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}

	Println(check(a))
}

func check(a []int) string {
	i1 := -1
	i2 := -1
	i3 := -1
	i4 := -1

	i := 1

	for i < len(a) {
		if a[i-1] > a[i] {
			i1 = i - 1
			break
		}
		i++
	}

	if i1 == -1 {
		return "yes"
	}

	if i <= len(a) {
		i = i1 + 1
	}

	for i < len(a) {
		if a[i-1] < a[i] {
			i2 = i - 1
			break
		}
		i++
	}

	if i == len(a) {
		i2 = len(a) - 1
	}

	for i < len(a) {
		if a[i-1] > a[i] {
			i3 = i - 1
			break
		}
		i++
	}

	if i3 == -1 {
		if trySwap(a, i1, i2) {
			return Sprintf("yes\nswap %v %v", i1+1, i2+1)
		}
		if tryReverse(a, i1, i2) {
			return Sprintf("yes\nreverse %v %v", i1+1, i2+1)
		}
		return "no"
	}

	if i <= len(a) {
		i = i3 + 1
	}

	for i < len(a) {
		if a[i-1] < a[i] {
			i4 = i - 1
			break
		}
		i++
	}

	if i == len(a) {
		i4 = len(a) - 1
	}

	if i2-i1 == 1 && i4-i3 == 1 && trySwap(a, i1, i4) {
		return Sprintf("yes\nswap %v %v", i1+1, i4+1)
	}

	return "no"
}

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}

func trySwap(a []int, i, j int) bool {
	a[i], a[j] = a[j], a[i]
	v := isSorted(a)
	a[i], a[j] = a[j], a[i]
	return v
}

func tryReverse(a []int, i, j int) bool {
	for q := 0; q <= (j-i)/2; q++ {
		a[i+q], a[j-q] = a[j-q], a[i+q]
	}
	v := isSorted(a)
	for q := 0; q <= (j-i)/2; q++ {
		a[i+q], a[j-q] = a[j-q], a[i+q]
	}
	return v
}
