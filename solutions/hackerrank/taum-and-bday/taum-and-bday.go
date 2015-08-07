package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/taum-and-bday

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	for i := 0; i < n; i++ {
		s.Scan()
		b, _ := strconv.Atoi(s.Text())
		s.Scan()
		w, _ := strconv.Atoi(s.Text())
		s.Scan()
		x, _ := strconv.Atoi(s.Text())
		s.Scan()
		y, _ := strconv.Atoi(s.Text())
		s.Scan()
		z, _ := strconv.Atoi(s.Text())

		// cases:
		// * x y z
		// * x z y
		// * y x z
		// * y z x
		// * z x y
		// * z y x
		// * x=y z
		// * z x=y
		// * x=z y
		// * y x=z
		// * y=z x
		// * x y=z => b*x+w*y
		// * x=y=z => b*x+w*y

		// formulas:
		// buy all black and convert: b*x + w*(x+z)
		// buy all white and convert: b*(y+z) + w*y
		// buy black and white: b*x+w*y

		s1 := b*x + w*(x+z)
		s2 := b*(y+z) + w*y
		s3 := b*x + w*y

		var m12 int
		if s1 > s2 {
			m12 = s2
		} else {
			m12 = s1
		}
		if m12 > s3 {
			fmt.Printf("%v\n", s3)
		} else {
			fmt.Printf("%v\n", m12)
		}
	}
}
