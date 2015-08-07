package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/alternating-characters

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadBytes('\n')
	n, _ := strconv.Atoi(strings.Trim(string(line), "\n"))

	for i := 0; i < n; i++ {
		line, _ := r.ReadBytes('\n')
		cnt := 0
		pc := ' '
		for i, c := range strings.Trim(string(line), "\n") {
			if i > 0 && c == pc {
				cnt++
			}
			pc = c
		}

		fmt.Printf("%v\n", cnt)
	}
}
