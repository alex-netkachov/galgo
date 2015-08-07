package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/sherlock-and-squares

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	for i := 0; i < n; i++ {
		s.Scan()
		a, _ := strconv.Atoi(s.Text())
		s.Scan()
		b, _ := strconv.Atoi(s.Text())

		cnt := 0
		start := int(math.Floor(math.Sqrt(float64(a))))
		finish := int(math.Ceil(math.Sqrt(float64(b))))
		for i := start; i <= finish; i++ {
			v := i * i
			if v >= a && v <= b {
				cnt++
			}
		}
		fmt.Printf("%v\n", cnt)
	}
}
