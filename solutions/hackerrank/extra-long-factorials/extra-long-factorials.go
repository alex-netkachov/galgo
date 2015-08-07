package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/extra-long-factorials

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	f := big.NewInt(int64(1))
	for i := 1; i <= n; i++ {
		f.Mul(f, big.NewInt(int64(i)))
	}

	fmt.Printf("%v\n", f)
}
