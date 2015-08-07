package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://www.hackerrank.com/challenges/game-of-thrones

func main() {
	r := bufio.NewReader(os.Stdin)
	data, _ := r.ReadBytes('\n')
	line := strings.Trim(string(data), "\n")

	cs := make(map[rune]int)

	for _, c := range line {
		v, _ := cs[c]
		cs[c] = v + 1
	}

	odds := 0
	for _, v := range cs {
		if v%2 == 1 {
			odds++
		}
	}

	if len(line)%2 == odds {
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO\n")
	}
}
