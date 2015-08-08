package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt quicksort1.go && go run quicksort1.go < quicksort1.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Scan()
	p, _ := strconv.Atoi(s.Text())

	al := []int{}
	ar := []int{p}

	for i := 1; i < n; i++ {
		s.Scan()
		v, _ := strconv.Atoi(s.Text())
		if v > p {
			ar = append(ar, v)
		} else {
			al = append(al, v)
		}
	}

	for _, v := range append(al, ar...) {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
}
