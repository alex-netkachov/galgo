package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt insertionsort2.go && go run insertionsort2.go < insertionsort2.txt

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

	sort(&a)
}

func sort(array *[]int) {
	a := *array
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
			print(a)
			continue
		}
		v := a[i+1]
		j := i
		for ; j > -1; j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = v
		print(a)
	}
}

func print(a []int) {
	for _, v := range a {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
}
