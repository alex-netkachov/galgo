package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt connected-cell-in-a-grid.go && go run connected-cell-in-a-grid.go < connected-cell-in-a-grid.txt

type p struct {
	i, j int
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	m, _ := strconv.Atoi(s.Text())
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			s.Scan()
			a[i][j], _ = strconv.Atoi(s.Text())
		}
	}

	max := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				v := mark(a, i, j, m, n)
				if max < v {
					max = v
				}
			}
		}
	}

	Println(max)
}

func mark(a [][]int, i, j, m, n int) int {
	v := visit([]p{}, a, i, j, m, n)
	vi := 0

	for {
		q := v[vi]

		v = visit(v, a, q.i-1, q.j-1, m, n)
		v = visit(v, a, q.i-1, q.j, m, n)
		v = visit(v, a, q.i-1, q.j+1, m, n)

		v = visit(v, a, q.i, q.j+1, m, n)

		v = visit(v, a, q.i+1, q.j+1, m, n)
		v = visit(v, a, q.i+1, q.j, m, n)
		v = visit(v, a, q.i+1, q.j-1, m, n)

		v = visit(v, a, q.i, q.j-1, m, n)

		vi++
		if vi == len(v) {
			break
		}
	}

	return len(v)
}

func visit(v []p, a [][]int, i, j, m, n int) []p {
	if 0 <= i && i < m && 0 <= j && j < n && a[i][j] == 1 {
		a[i][j] = 2
		v = append(v, p{i, j})
	}
	return v
}
