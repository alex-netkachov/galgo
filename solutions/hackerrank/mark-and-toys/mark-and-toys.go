package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

// go fmt mark-and-toys.go && go run mark-and-toys.go < mark-and-toys.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	t := make([]int, n)

	for i := 0; i < n; i++ {
		s.Scan()
		t[i], _ = strconv.Atoi(s.Text())
	}

	sort.Ints(t)

	m := k
	j := 0
	for m > 0 && j < len(t) {
		m -= t[j]
		j++
	}

	Println(j - 1)
}
