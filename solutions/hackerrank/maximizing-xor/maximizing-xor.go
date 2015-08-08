package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt maximizing-xor.go && go run maximizing-xor.go < maximizing-xor.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	l, _ := strconv.Atoi(s.Text())
	s.Scan()
	r, _ := strconv.Atoi(s.Text())

	m := 0
	for a := l; a < r; a++ {
		for b := l + 1; b <= r; b++ {
			if a^b > m {
				m = a ^ b
			}
		}
	}

	Println(m)
}
