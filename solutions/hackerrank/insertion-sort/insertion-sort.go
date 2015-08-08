package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt insertion-sort.go && go run insertion-sort.go < insertion-sort.txt

// based on http://www.geeksforgeeks.org/avl-tree-set-1-insertion/

type Node struct {
	Value int
	Left  *Node
	Right *Node

	height int
	count  int
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func count(n *Node) int {
	if n == nil {
		return 0
	}
	return n.count
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
       y                               x
      / \     Right Rotation          / \
     x  T3    – – – – – – – >        T1  y
    / \       < - - - - - - -           / \
   T1  T2     Left Rotation           T2  T3
*/

func rotateLeft(x *Node) *Node {
	y := x.Right
	t2 := y.Left
	y.Left = x
	x.Right = t2
	x.height = max(height(x.Left), height(x.Right)) + 1
	y.height = max(height(y.Left), height(y.Right)) + 1

	cx := count(x) - count(x.Left) - count(y)
	cy := count(y) - count(x.Right) - count(y.Right)

	x.count = cx + count(x.Left) + count(x.Right)
	y.count = cy + count(y.Left) + count(y.Right)

	return y
}

func rotateRight(y *Node) *Node {
	x := y.Left
	T2 := x.Right
	x.Right = y
	y.Left = T2
	y.height = max(height(y.Left), height(y.Right)) + 1
	x.height = max(height(x.Left), height(x.Right)) + 1

	cx := count(x) - count(x.Left) - count(y.Left)
	cy := count(y) - count(x) - count(y.Right)
	y.count = cy + count(y.Left) + count(y.Right)
	x.count = cx + count(x.Left) + count(x.Right)

	return x
}

func balance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

func Insert(n *Node, value int) *Node {
	// 1. Perform the normal BST rotation
	if n == nil {
		return &Node{value, nil, nil, 1, 1}
	}

	n.count++

	if value < n.Value {
		n.Left = Insert(n.Left, value)
	} else if value > n.Value {
		n.Right = Insert(n.Right, value)
	} else if value == n.Value {
		return n
	}

	// 2. Update height of this ancestor node
	n.height = max(height(n.Left), height(n.Right)) + 1

	// 3. Get the balance factor of this ancestor node to check whether
	//    this node became unbalanced
	b := balance(n)

	// If this node becomes unbalanced, then there are 4 cases

	// Left Left Case
	if b > 1 && value < n.Left.Value {
		return rotateRight(n)
	}

	// Right Right Case
	if b < -1 && value > n.Right.Value {
		return rotateLeft(n)
	}

	// Left Right Case
	if b > 1 && value > n.Left.Value {
		n.Left = rotateLeft(n.Left)
		return rotateRight(n)
	}

	// Right Left Case
	if b < -1 && value < n.Right.Value {
		n.Right = rotateRight(n.Right)
		return rotateLeft(n)
	}

	// return the (unchanged) node pointer
	return n
}

func rank(n *Node, v int) int {
	c := n
	s := -1
	for {
		if c == nil {
			return -1
		}
		if v < c.Value {
			c = c.Left
			continue
		}
		if v > c.Value {
			s += c.count - count(c.Right)
			c = c.Right
			continue
		}
		if v == c.Value {
			s += c.count - count(c.Right)
			return s
		}
	}
	return -1
}

func shifts(a []int) int {
	n := len(a)

	s := 0

	var r *Node
	for j := 0; j < n; j++ {
		r = Insert(r, a[j])
		s += j - rank(r, a[j])
	}

	return s
}

func main() {
	//test()

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
			v, _ := strconv.Atoi(s.Text())
			a[j] = v
		}

		Println(shifts(a))
	}
}
