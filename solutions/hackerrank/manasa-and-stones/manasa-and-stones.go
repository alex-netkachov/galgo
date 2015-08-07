package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/manasa-and-stones

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	for i := 0; i < n; i++ {
		s.Scan()
		ns, _ := strconv.Atoi(s.Text())
		s.Scan()
		a, _ := strconv.Atoi(s.Text())
		s.Scan()
		b, _ := strconv.Atoi(s.Text())

		if a > b {
			t := b
			b = a
			a = t
		}

		d := b - a
		if d == 0 {
			fmt.Printf("%v", a*(ns-1))
		} else {
			for c := a * (ns - 1); c <= b*(ns-1); c += (b - a) {
				fmt.Printf("%v ", c)
			}
		}

		fmt.Printf("\n")
	}
}
