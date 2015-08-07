package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// go fmt extra-long-factorials.go && go run extra-long-factorials.go < extra-long-factorials.txt

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
