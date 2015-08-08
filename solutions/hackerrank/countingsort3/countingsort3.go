package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt countingsort3.go && go run countingsort3.go < countingsort3.txt

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

		s.Scan()
	}

	sum := 0
	for i := 0; i < 100; i++ {
		sum += m[i]
		fmt.Printf("%v ", sum)
	}
	fmt.Printf("\n")
}
