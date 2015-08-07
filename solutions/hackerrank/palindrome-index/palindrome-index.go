package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	d, _ := r.ReadBytes('\n')
	n, _ := strconv.Atoi(strings.Trim(string(d), "\n"))

	for i := 0; i < n; i++ {
		d, _ = r.ReadBytes('\n')
		line := strings.Trim(string(d), "\n")

		if check(line, -1) {
			fmt.Printf("-1\n")
			continue
		}

		for j := 0; j < len(line); j++ {
			if check(line, j) {
				fmt.Printf("%v\n", j)
				break
			}
		}
	}
}

func check(word string, ex int) bool {
	i := -1
	j := len(word)
	for {
		i += 1
		j -= 1
		if i == ex {
			i += 1
		}
		if j == ex {
			j -= 1
		}
		if i > j {
			return true
		}
		if word[i] != word[j] {
			return false
		}
	}
	return true
}
