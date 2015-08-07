package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	ss := " "
	sh := "#"
	for len(ss) < n {
		ss += ss
		sh += sh
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%v%v\n", ss[0:n-i], sh[0:i])
	}
}
