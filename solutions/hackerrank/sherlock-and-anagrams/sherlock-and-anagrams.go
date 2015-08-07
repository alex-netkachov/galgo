package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// https://www.hackerrank.com/challenges/sherlock-and-anagrams

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanLines)

	for i := 0; i < n; i++ {
		s.Scan()
		line := s.Text()

		a := make(map[string]int)

		for j := 0; j < len(line); j++ {
			for k := j + 1; k < len(line)+1; k++ {
				if j == 0 && k == len(line) {
					continue
				}

				f := base(line[j:k])

				v := a[f]
				a[f] = v + 1
			}
		}

		cnt := 0
		for _, v := range a {
			cnt += v * (v - 1)
		}

		fmt.Printf("%v\n", cnt/2)
	}
}

func base(str string) string {
	b := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		b[i] = int(str[i])
	}
	sort.Ints(b)
	t := make([]byte, len(str))
	for i := 0; i < len(b); i++ {
		t[i] = byte(b[i])
	}
	return string(t)
}
