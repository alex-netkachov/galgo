package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// https://www.hackerrank.com/challenges/encryption

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))

	s.Scan()
	text := s.Text()

	cs := make([]string, int(math.Ceil(math.Sqrt(float64(len(text))))))

	for i := 0; i < len(text); i++ {
		cs[i%len(cs)] += text[i : i+1]
	}

	for _, v := range cs {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
}
