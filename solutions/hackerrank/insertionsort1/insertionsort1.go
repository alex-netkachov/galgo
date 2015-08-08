package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt insertionsort1.go && go run insertionsort1.go < insertionsort1.txt

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

	for i := 0; i < n-1; i++ {
		if a[i] < a[i+1] {
			continue
		}
		v := a[i+1]
		j := i
		for ; j > -1; j-- {
			if a[j] > v {
				a[j+1] = a[j]
				print(a)
			} else {
				break
			}
		}
		a[j+1] = v
		print(a)
		break
	}
}

func print(a []int) {
	for _, v := range a {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
}
