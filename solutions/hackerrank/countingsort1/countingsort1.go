package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt countingsort1.go && go run countingsort1.go < countingsort1.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		s.Scan()
		v, _ := strconv.Atoi(s.Text())
		c, _ := m[v]
		m[v] = c + 1
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("%v ", m[i])
	}
	fmt.Printf("\n")
}
