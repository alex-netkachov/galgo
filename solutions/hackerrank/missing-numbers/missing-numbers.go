package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

// go fmt missing-numbers.go && go run missing-numbers.go < missing-numbers.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	a := ReadInts(s)
	b := ReadInts(s)

	ma := make(map[int]int)
	mb := make(map[int]int)

	for _, v := range a {
		i, _ := ma[v]
		ma[v] = i + 1
	}
	for _, v := range b {
		i, _ := mb[v]
		mb[v] = i + 1
	}

	keys := func(m map[int]int) []int {
		k := make([]int, len(m))
		i := 0
		for mk, _ := range m {
			k[i] = mk
			i++
		}
		sort.Ints(k)
		return k
	}

	ka := keys(ma)
	kb := keys(mb)

	ia := 0
	ib := 0
	for {
		if ia == len(ka) && ib == len(kb) {
			break
		}
		if ia == len(ka) {
			k := kb[ib]
			Printf("%v ", k)
			ib++
			continue
		}
		if ib == len(kb) {
			k := ka[ia]
			Printf("%v ", k)
			ia++
			continue
		}
		kia := ka[ia]
		kib := kb[ib]
		if kia < kib {
			Printf("%v ", kia)
			ia++
			continue
		}
		if kia > kib {
			Printf("%v ", kib)
			ib++
			continue
		}
		if kia == kib {
			if mb[kib] != ma[kia] {
				Printf("%v ", kia)
			}
			ia++
			ib++
			continue
		}
	}
	Println()
}

func ReadInts(s *bufio.Scanner) []int {
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}
	return a
}
