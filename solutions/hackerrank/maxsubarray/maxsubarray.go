package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt maxsubarray1.go && go run maxsubarray1.go < maxsubarray.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())

		a := make([]int, n)
		for j := 0; j < n; j++ {
			s.Scan()
			a[j], _ = strconv.Atoi(s.Text())
		}

		max := a[0]
		for _, v := range a {
			if max < v {
				max = v
			}
		}

		if max <= 0 {
			Println(max, max)
			continue
		}

		c := compact(a)
		m1, m2 := do(c)
		Println(m1, m2)
	}
}

func do(a []int) (int, int) {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}

	bp := make([]int, len(a))
	copy(bp, a)
	bc := make([]int, len(a))

	a2 := make([]int, len(a)-1)
	for i := 0; i < len(a2); i++ {
		a2[i] = a[i] + a[i+1]
	}

	for i := len(a) - 2; i > 0; i -= 2 {
		d := len(a) - i - 1
		for j, v := range bp[0:i] {
			c := v + a2[j+d]
			bc[j] = c
			if c > max {
				max = c
			}
		}
		bp, bc = bc, bp
	}

	s := 0
	for _, v := range a {
		if v > 0 {
			s += v
		}
	}

	return max, s
}

func compact(a []int) []int {
	var g []int
	var s int
	if a[0] > 0 {
		g = []int{a[0]}
		s = 1
	} else {
		g = []int{0, a[0]}
		s = -1
	}

	for j := 1; j < len(a); j++ {
		v := a[j]
		if s > 0 {
			if v >= 0 {
				g[len(g)-1] += v
				continue
			}
			s = -1
			g = append(g, v)
		} else {
			if v <= 0 {
				g[len(g)-1] += v
				continue
			}
			s = 1
			g = append(g, v)
		}
	}

	if s == 1 {
		g = append(g, 0)
	}

	return g
}

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}
