package main

import (
	. "fmt"
)

// go fmt count-luck.go && go run count-luck.go < count-luck.txt

const cm byte = byte('M')
const ct byte = byte('X')
const ce byte = byte('*')
const cs byte = byte('.')

type pt struct{ i, j int }

func main() {
	t := 0
	Scanln(&t)

	for i := 0; i < t; i++ {
		n := 0
		Scan(&n)

		m := 0
		Scan(&m)

		f := make([][]byte, n)
		for j := 0; j < n; j++ {
			f[j] = make([]byte, m)

			s := ""
			Scanln(&s)
			for k, v := range s {
				f[j][k] = byte(v)
			}
		}

		k := 0
		Scan(&k)

		kc := 0
		path := findpath(f)
		wp := []pt{}
		for i, p := range path {
			wc := 0
			for _, q := range getinc(f, p) {
				if f[q.i][q.j] != ct {
					wc++
				}
			}
			if i == 0 {
				if wc > 1 {
					wp = append(wp, p)
					kc++
				}
			} else {
				if wc > 2 {
					wp = append(wp, p)
					kc++
				}
			}
		}

		if k == kc {
			Println("Impressed")
		} else {
			Println("Oops!")
		}
	}
}

func findpath(f [][]byte) []pt {
	m := make([][]int, len(f))
	for i := 0; i < len(f); i++ {
		m[i] = make([]int, len(f[i]))
	}

	hi, hj := findvalue(f, cm)
	ei, ej := findvalue(f, ce)

	// propagate
	vs := []pt{pt{hi, hj}}
	m[hi][hj] = 1
	for i := 0; i < len(vs); i++ {
		v := vs[i]
		if v.i == ei && v.j == ej {
			break
		}
		s := m[v.i][v.j] + 1
		for _, p := range getinc(f, v) {
			if f[p.i][p.j] == cs || f[p.i][p.j] == ce {
				if m[p.i][p.j] == 0 {
					m[p.i][p.j] = s
					vs = append(vs, p)
				}
			}
		}
		//printfieldm()
	}

	// fold
	q := pt{ei, ej}
	r := []pt{}
	for v := m[q.i][q.j] - 1; v > 0; v-- {
		for _, p := range getinc(f, q) {
			if m[p.i][p.j] == v {
				r = append(r, p)
				q = p
				break
			}
		}
	}

	// reverse
	p := make([]pt, len(r))
	for i := 0; i < len(r); i++ {
		p[len(p)-1-i] = r[i]
	}

	return p
}

func getinc(f [][]byte, p pt) []pt {
	r := make([]pt, 0)
	if p.i > 0 {
		r = append(r, pt{p.i - 1, p.j})
	}
	if p.i < len(f)-1 {
		r = append(r, pt{p.i + 1, p.j})
	}
	if p.j > 0 {
		r = append(r, pt{p.i, p.j - 1})
	}
	if p.j < len(f[p.i])-1 {
		r = append(r, pt{p.i, p.j + 1})
	}
	return r
}

func findvalue(f [][]byte, v byte) (int, int) {
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			if f[i][j] == v {
				return i, j
			}
		}
	}
	return -1, -1
}

func printfield(f [][]byte) {
	for _, r := range f {
		Println(string(r))
	}
}
