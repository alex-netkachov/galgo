package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		t, _ := strconv.ParseUint(s.Text(), 10, 32)
		ut := uint32(t)
		fmt.Printf("%v\n", ^ut)
	}
}
