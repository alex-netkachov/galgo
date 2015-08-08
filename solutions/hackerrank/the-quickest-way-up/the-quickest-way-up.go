package main

import (
	. "fmt"
)

// go fmt the-quickest-way-up.go && go run the-quickest-way-up.go < the-quickest-way-up.txt

type node struct {
	i int
	j []*node
}

func makeNode(i int) *node {
	return &node{i, []*node{nil, nil, nil, nil, nil, nil}}
}

func main() {
	t := 0
	Scanln(&t)

	for i := 0; i < t; i++ {
		field := make([]int, 100)
		for i := 0; i < 100; i++ {
			field[i] = i + 1
		}

		sn := 0
		Scan(&sn)

		for j := 0; j < sn; j++ {
			from := 0
			Scan(&from)
			to := 0
			Scan(&to)
			field[from-1] = to - 1
		}

		ln := 0
		Scan(&ln)

		for j := 0; j < ln; j++ {
			from := 0
			Scan(&from)
			to := 0
			Scan(&to)
			field[from-1] = to - 1
		}

		nodes := make([]*node, 100)
		for i := 0; i < 100; i++ {
			nodes[i] = makeNode(i)
		}
		for i := 0; i < 100; i++ {
			for j := 0; j < 6; j++ {
				p := i + j + 1
				if p < 100 {
					if field[p] != p+1 {
						p = field[p]
					}
					nodes[i].j[j] = nodes[p]
				}
			}
		}

		visited := make([]int, 100)
		run := []*node{nodes[0]}

		for i := 0; i < len(run); i++ {
			for _, n := range run[i].j {
				if n != nil && visited[n.i] == 0 {
					visited[n.i] = visited[run[i].i] + 1
					run = append(run, n)
				}
			}
		}

		if visited[99] == 0 {
			Println("-1")
		} else {
			Println(visited[99])
		}
	}
}
