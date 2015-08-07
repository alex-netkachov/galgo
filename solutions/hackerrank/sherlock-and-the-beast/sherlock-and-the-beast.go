package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/sherlock-and-the-beast

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		i3 := 0
		i5 := 0

		for ; i3 < n; i3 += 5 {
			v := (n - i3) / 3
			if v*3 == n-i3 {
				i5 = n - i3
				break
			}
		}

		if (i3 + i5) == n {
			var s5 bytes.Buffer
			for j := 0; j < i5; j++ {
				s5.WriteString("5")
			}
			var s3 bytes.Buffer
			for j := 0; j < i3; j++ {
				s3.WriteString("3")
			}
			fmt.Printf("%s%s\n", s5.String(), s3.String())
		} else {
			fmt.Printf("-1\n")
		}
	}
}
