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
	n, _ := strconv.Atoi(s.Text())

	s.Split(bufio.ScanLines)

	s.Scan()
	t := s.Text()
	t = t[0:n]

	s.Split(bufio.ScanWords)

	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	cA := int("A"[0])
	cZ := int("Z"[0]) + 1
	ca := int("a"[0])
	cz := int("z"[0]) + 1

	r := ""
	for i := 0; i < len(t); i++ {
		v := int(t[i])
		if v >= cA && v < cZ {
			r += string(((v-cA)+(cZ-cA)+k)%(cZ-cA) + cA)
		} else if v >= ca && v < cz {
			r += string(((v-ca)+(cz-ca)+k)%(cz-ca) + ca)
		} else {
			r += string(v)
		}
	}

	fmt.Printf("%v\n", r)
}
