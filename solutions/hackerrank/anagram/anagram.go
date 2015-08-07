package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/anagram

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanLines)

	for i := 0; i < n; i++ {
		s.Scan()
		l := s.Text()

		if len(l)%2 == 1 {
			fmt.Printf("-1\n")
			continue
		}

		cs1 := make(map[rune]int)
		cs2 := make(map[rune]int)

		for _, c := range l[0 : len(l)/2] {
			v, _ := cs1[c]
			cs1[c] = v + 1
		}

		for _, c := range l[len(l)/2:] {
			v, _ := cs2[c]
			cs2[c] = v + 1
		}

		add := 0
		remove := 0
		for c1, v1 := range cs1 {
			v2, ok := cs2[c1]
			if ok {
				v := v1 - v2
				if v > 0 {
					remove += v
				} else {
					add += -v
				}
			} else {
				remove += v1
			}
		}

		fmt.Printf("%v\n", remove)
	}

}
