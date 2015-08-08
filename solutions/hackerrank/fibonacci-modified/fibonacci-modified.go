package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// go fmt fibonacci-modified.go && go run fibonacci-modified.go < fibonacci-modified.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	a, _ := strconv.Atoi(s.Text())

	s.Scan()
	b, _ := strconv.Atoi(s.Text())

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	switch n {
	case 0:
		fmt.Printf("0\n")
	case 1:
		fmt.Printf("%v\n", a)
	case 2:
		fmt.Printf("%v\n", b)
	default:
		fp := big.NewInt(int64(a))
		fc := big.NewInt(int64(b))
		for i := 2; i < n; i++ {
			f := new(big.Int)
			f.Mul(fc, fc)
			f.Add(f, fp)
			fp = fc
			fc = f
		}
		fmt.Printf("%v\n", fc)
	}
}
