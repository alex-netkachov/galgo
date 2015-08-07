package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/bigger-is-greater

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanLines)

	for i := 0; i < n; i++ {
		s.Scan()
		line := s.Text()

		a := make([]byte, len(line))
		for j := 0; j < len(line); j++ {
			a[j] = line[j]
		}

		found := false
		for j := len(a) - 2; j > -1; j-- {
			k := j
			v := a[k]
			if v >= a[len(a)-1] {
				for ; k < len(a)-1; k++ {
					a[k] = a[k+1]
				}
				a[len(a)-1] = v
			} else {
				for k++; k < len(a) && v >= a[k]; k++ {
				}
				a[j], a[k] = a[k], v
				found = string(a) != line
				break
			}
		}
		if found {
			fmt.Printf("%v\n", string(a))
		} else {
			fmt.Printf("no answer\n")
		}
	}
}
