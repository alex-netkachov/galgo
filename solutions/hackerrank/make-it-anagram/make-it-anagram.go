package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.hackerrank.com/challenges/make-it-anagram

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))

	s.Scan()
	l1 := s.Text()
	s.Scan()
	l2 := s.Text()

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

	cnt := 0
	for c1, v1 := range cs1 {
		v2, ok := cs2[c1]
		if ok {
			v := v1 - v2
			if v > 0 {
				cnt += v
			} else {
				cnt += -v
			}
		} else {
			cnt += v1
		}
	}

	for c2, v2 := range cs2 {
		_, ok := cs1[c2]
		if !ok {
			cnt += v2
		}
	}

	fmt.Printf("%v\n", cnt)
}
