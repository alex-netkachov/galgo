package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt lonely-integer.go && go run lonely-integer.go < lonely-integer.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	c := make([]int, 101)
	for i := 0; i < n; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		c[n]++
	}
	for i, v := range c {
		if v == 1 {
			Println(i)
		}
	}
}
