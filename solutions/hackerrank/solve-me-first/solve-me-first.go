package main

import (
	"fmt"
)

// https://www.hackerrank.com/challenges/solve-me-first

func solveMeFirst(a uint32, b uint32) uint32 {
	return a + b
}

func main() {
	var a, b uint32

	fmt.Scanf("%v %v", &a, &b)

	res := solveMeFirst(a, b)

	fmt.Println(res)
}
