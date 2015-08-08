package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// go fmt angry-children.go && go run angry-children.go < angry-children.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	ints := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		ints[i], _ = strconv.Atoi(s.Text())
	}

	sort.Ints(ints)

	min := ints[k-1] - ints[0]
	for i := 1; i < n-k+1; i++ {
		v := ints[i+k-1] - ints[i]
		if v < min {
			min = v
		}
	}

	fmt.Printf("%v\n", min)
}
