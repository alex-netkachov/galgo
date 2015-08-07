package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://www.hackerrank.com/challenges/pangrams

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	line := strings.ToLower(s.Text())

	pangram := true
	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		if !strings.Contains(line, string(c)) {
			pangram = false
		}
	}

	if pangram {
		fmt.Printf("pangram\n")
	} else {
		fmt.Printf("not pangram\n")
	}
}
