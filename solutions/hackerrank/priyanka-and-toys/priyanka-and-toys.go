package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

// go fmt priyanka-and-toys.go && go run priyanka-and-toys.go < priyanka-and-toys.txt

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

	sort.Ints(a)

	q := 0
	for {
		c := 0
		for _, v := range a {
			if a[0] <= v && v <= a[0]+4 {
				c++
			} else {
				break
			}
		}
		q++
		if c == len(a) {
			break
		}
		a = a[c:]
	}

	Println(q)
}
