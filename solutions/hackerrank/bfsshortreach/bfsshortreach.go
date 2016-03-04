package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/bfsshortreach

type Node struct {
	Nodes []int
	Path  int
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	t := ReadInt(s)

	for i := 0; i < t; i++ {
		n := ReadInt(s)
		m := ReadInt(s)

		g := []Node{}
		for j := 0; j < n; j++ {
			g = append(g, Node{[]int{}, -1})
		}

		for j := 0; j < m; j++ {
			x := ReadInt(s) - 1
			y := ReadInt(s) - 1

			g[x].Nodes = append(g[x].Nodes, y)
			g[y].Nodes = append(g[y].Nodes, x)
		}

		st := ReadInt(s) - 1

		PrintDistances(g, st)
	}
}

func PrintDistances(g []Node, start int) {
	g[start].Path = 0

	queue := []int{}
	queue = append(queue, start)
	index := 0
	for index < len(queue) {
		for _, j := range g[queue[index]].Nodes {
			if g[j].Path < 0 {
				queue = append(queue, j)
				g[j].Path = g[queue[index]].Path + 6
			}
		}
		index++
	}

	for _, n := range g {
		if n.Path != 0 {
			fmt.Printf("%d ", n.Path)
		}
	}
	fmt.Println("")
}

func ReadInt(s *bufio.Scanner) int {
	// call s.Split(bufio.ScanWords) before calling this method

	s.Scan()
	i, _ := strconv.Atoi(s.Text())

	return i
}
