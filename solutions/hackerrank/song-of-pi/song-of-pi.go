package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/song-of-pi

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	pi := "31415926535897932384626433833"

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanLines)

	for i := 0; i < n; i++ {
		s.Scan()
		words := strings.Split(s.Text(), " ")

		piStr := ""
		for _, w := range words {
			piStr = fmt.Sprintf("%s%d", piStr, len(w))
		}
		if 0 == strings.Index(pi, piStr) {
			fmt.Printf("It's a pi song.\n")
		} else {
			fmt.Printf("It's not a pi song.\n")
		}
	}
}
