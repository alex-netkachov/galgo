package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// go fmt icecream-parlor.go && go run icecream-parlor.go < icecream-parlor.txt

type flavor struct {
	price, index int
}

type byprice []flavor

func (a byprice) Len() int           { return len(a) }
func (a byprice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byprice) Less(i, j int) bool { return a[i].price < a[j].price }

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		m, _ := strconv.Atoi(s.Text())

		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		c := make([]flavor, n)
		for j := 0; j < n; j++ {
			s.Scan()
			v, _ := strconv.Atoi(s.Text())
			c[j] = flavor{v, j + 1}
		}
		sort.Sort(byprice(c))

		for j := 0; j < n; j++ {
			v := c[j].price
			q := sort.Search(n-j-1, func(i int) bool {
				return c[i+j+1].price+v >= m
			}) + j + 1

			if q < len(c) && c[q].price+v == m {
				if c[q].index > c[j].index {
					fmt.Printf("%v %v\n", c[j].index, c[q].index)
				} else {
					fmt.Printf("%v %v\n", c[q].index, c[j].index)
				}
				break
			}
		}
	}
}
