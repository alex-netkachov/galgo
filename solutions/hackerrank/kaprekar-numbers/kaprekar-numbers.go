package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/kaprekar-numbers

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	p, _ := strconv.Atoi(s.Text())

	s.Scan()
	q, _ := strconv.Atoi(s.Text())

	numbers := make([]int, 0)
	for i := p; i <= q; i++ {
		if isKaprekar(i) {
			numbers = append(numbers, i)
		}
	}

	if len(numbers) == 0 {
		fmt.Printf("INVALID RANGE\n")
	} else {
		for _, n := range numbers {
			fmt.Printf("%v ", n)
		}
		fmt.Printf("\n")
	}
}

func isKaprekar(n int) bool {
	d := len(digitise(n))
	dg := digitise(n * n)
	l := undigitise(dg[0:d])
	r := undigitise(dg[d:])
	if n == l+r {
		return true
	}
	return false
}

func digitise(n int) []int {
	r := make([]int, 0)
	for ; n > 0; n /= 10 {
		r = append(r, n%10)
	}
	return r
}

func undigitise(a []int) int {
	v := 0
	for i := len(a) - 1; i >= 0; i-- {
		v = v*10 + a[i]
	}
	return v
}
