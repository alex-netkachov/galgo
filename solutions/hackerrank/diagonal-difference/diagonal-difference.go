package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s.Scan()
			v, _ := strconv.Atoi(s.Text())
			m[i][j] = v
		}
	}

	sum1 := 0
	sum2 := 0
	for i := 0; i < n; i++ {
		sum1 += m[i][i]
		sum2 += m[i][n-i-1]
	}

	fmt.Printf("%v\n", math.Abs(float64(sum1-sum2)))
}
