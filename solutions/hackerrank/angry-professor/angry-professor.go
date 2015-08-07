package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/angry-professor

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	for i := 0; i < n; i++ {
		s.Scan()
		tn, _ := strconv.Atoi(s.Text())
		s.Scan()
		tk, _ := strconv.Atoi(s.Text())
		cnt := 0
		for j := 0; j < tn; j++ {
			s.Scan()
			t, _ := strconv.Atoi(s.Text())
			if t > 0 {
				cnt++
			}
		}
		if (tn - cnt) < tk {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	}
}
