package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func pad2(v int) string {
	if v < 10 {
		return fmt.Sprintf("0%v", v)
	}
	return fmt.Sprintf("%v", v)
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanLines)

	s.Scan()
	ts := s.Text()

	h, _ := strconv.Atoi(ts[0:2])
	m, _ := strconv.Atoi(ts[3:5])
	sc, _ := strconv.Atoi(ts[6:8])
	am := ts[8:10]

	if am == "AM" {
		if h == 12 {
			h = 0
		}
	} else {
		if h != 12 {
			h += 12
		}
	}

	fmt.Printf("%v:%v:%v\n", pad2(h), pad2(m), pad2(sc))
}
