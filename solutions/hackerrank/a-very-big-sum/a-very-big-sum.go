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

	sum := uint64(0)
	for i := 0; i < n; i++ {
		s.Scan()
		v, _ := strconv.Atoi(s.Text())
		sum += uint64(v)
	}

	fmt.Printf("%v\n", sum)
}
