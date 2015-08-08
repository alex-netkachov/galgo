package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt a-very-big-sum.go && go run a-very-big-sum.go < a-very-big-sum.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		a := make(map[int]int)
		for j := 0; j < n; j++ {
			s.Scan()
			v, _ := strconv.Atoi(s.Text())
			c, _ := a[v]
			a[v] = c + 1
		}

		s := 0
		for _, v := range a {
			s += v * (v - 1)
		}
		Println(s)
	}
}
