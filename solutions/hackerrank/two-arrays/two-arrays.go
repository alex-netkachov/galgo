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
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		s.Scan()
		k, _ := strconv.Atoi(s.Text())

		a := make([]int, n)
		for j := 0; j < n; j++ {
			s.Scan()
			a[j], _ = strconv.Atoi(s.Text())
		}

		b := make([]int, n)
		for j := 0; j < n; j++ {
			s.Scan()
			b[j], _ = strconv.Atoi(s.Text())
		}

		sort.Ints(a)
		sort.Ints(b)

		ok := true
		for j := 0; j < n; j++ {
			if a[j]+b[n-1-j] < k {
				ok = false
				break
			}
		}

		if ok {
			Println("YES")
		} else {
			Println("NO")
		}
	}
}
