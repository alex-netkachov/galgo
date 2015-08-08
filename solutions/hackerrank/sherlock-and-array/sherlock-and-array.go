package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt sherlock-and-array.go && go run sherlock-and-array.go < sherlock-and-array.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		a := make([]int, n)
		for j := 0; j < n; j++ {
			s.Scan()
			a[j], _ = strconv.Atoi(s.Text())
		}

		p, q := 0, len(a)-1
		sp, sq := 0, 0
		for {
			if p == q {
				if sq == sp {
					fmt.Printf("YES\n")
				} else {
					fmt.Printf("NO\n")
				}
				break
			}
			if sp > sq {
				sq += a[q]
				q--
			} else {
				sp += a[p]
				p++
			}
		}
	}
}
