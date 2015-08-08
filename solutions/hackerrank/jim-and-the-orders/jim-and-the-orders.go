package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

// go fmt jim-and-the-orders.go && go run jim-and-the-orders.go < jim-and-the-orders.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	q := make(map[int][]int)

	for i := 0; i < n; i++ {
		s.Scan()
		t, _ := strconv.Atoi(s.Text())
		s.Scan()
		d, _ := strconv.Atoi(s.Text())

		r := t + d
		v, ok := q[r]
		if !ok {
			v = []int{i}
		} else {
			v = append(v, i)
		}
		q[r] = v
	}

	t := make([]int, len(q))
	i := 0
	for k, v := range q {
		t[i] = k
		i++
		sort.Ints(v)
	}
	sort.Ints(t)

	for _, i := range t {
		v := q[i]
		for _, f := range v {
			Printf("%v ", f+1)
		}
	}
	Println()
}
