package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/two-strings

func main() {
	r := bufio.NewReader(os.Stdin)

	d, _ := r.ReadBytes('\n')
	n, _ := strconv.Atoi(strings.Trim(string(d), "\n"))

	for i := 0; i < n; i++ {
		d, _ = r.ReadBytes('\n')
		l1 := strings.Trim(string(d), "\n")
		d, _ = r.ReadBytes('\n')
		l2 := strings.Trim(string(d), "\n")

		cs1 := make(map[rune]int)
		cs2 := make(map[rune]int)

		for _, c := range l1 {
			v, _ := cs1[c]
			cs1[c] = v + 1
		}

		for _, c := range l2 {
			v, _ := cs2[c]
			cs2[c] = v + 1
		}

		intersection := false
		for c1, _ := range cs1 {
			_, ok := cs2[c1]
			if ok {
				intersection = true
			}
		}

		if intersection {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	}

}
