package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt maximise-sum.go && go run maximise-sum.go < maximise-sum.txt

func main() {
	//test()

	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	t := ReadLong(s)

	for i := int64(0); i < t; i++ {
		n := ReadLong(s)
		m := ReadLong(s)

		a := make([]int64, n)
		for j := int64(0); j < n; j++ {
			a[j] = ReadLong(s)
		}

		Println(maxsum(m, a))
	}
}

func maxsum(m int64, a []int64) int64 {
	max := int64(0)

	n := len(a)

	// p[0] = a[0] % M
	// p[1] = (a[0] + a[1]) % M
	// p[2] = (a[0] + a[1] + a[2]) % M
	// ...
	// p[n] = (a[0] + a[1] + a[2] + ... + a[n]) % M

	p := make([]int64, n)
	s := int64(0)

	// p[3] % M = (a[0] + a[1] + a[2] + a[3]) % M
	// Loop:
	//  (p[3] - p[0]) % M = (a[0] + a[1] + a[2] + a[3]) % M - a[0] % M = (a[1] + a[2] + a[3]) % M
	//  (p[3] - p[1]) % M = (a[2] + a[3]) % M
	//  (p[3] - p[2]) % M = a[3] % M
	//
	// Optimization:
	//  in p[3] - p[2] if 0 <= p[2] <= p[3] then the result will be in [0;p[3]]
	//  so check only if p[2] > p[3]

	var r *Node

	for i := 0; i < n; i++ {
		s = (s + a[i]) % m
		p[i] = s

		r = InsertOrIgnore(r, s)

		if s > max {
			max = s
		}

		nn := Next(r, s)
		if nn != nil {
			d := (s - nn.Key + m) % m
			if d > max {
				max = d
			}
		}
	}

	return max
}

func ReadLong(s *bufio.Scanner) int64 {
	s.Scan()
	i, _ := strconv.ParseInt(s.Text(), 10, 64)
	return i
}

type Node struct {
	Key    int64
	Left   *Node
	Right  *Node
	height int
}

func createNode(key int64) *Node {
	n := &Node{}
	n.Key = key
	n.height = 1
	return n
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func rotateLeft(x *Node) *Node {
	y := x.Right
	t2 := y.Left
	y.Left = x
	x.Right = t2
	x.height = max(height(x.Left), height(x.Right)) + 1
	y.height = max(height(y.Left), height(y.Right)) + 1
	return y
}

func rotateRight(y *Node) *Node {
	x := y.Left
	t2 := x.Right
	x.Right = y
	y.Left = t2
	y.height = max(height(y.Left), height(y.Right)) + 1
	x.height = max(height(x.Left), height(x.Right)) + 1
	return x
}

func balance(n *Node, key int64) *Node {
	b := height(n.Left) - height(n.Right)

	if b > 1 && key < n.Left.Key {
		return rotateRight(n)
	}

	if b < -1 && key > n.Right.Key {
		return rotateLeft(n)
	}

	if b > 1 && key > n.Left.Key {
		n.Left = rotateLeft(n.Left)
		return rotateRight(n)
	}

	if b < -1 && key < n.Right.Key {
		n.Right = rotateRight(n.Right)
		return rotateLeft(n)
	}

	return n
}

func InsertOrIgnore(n *Node, key int64) *Node {
	if n == nil {
		return createNode(key)
	}

	if key < n.Key {
		n.Left = InsertOrIgnore(n.Left, key)
	} else if key > n.Key {
		n.Right = InsertOrIgnore(n.Right, key)
	} else if key == n.Key {
		return n
	}

	n.height = max(height(n.Left), height(n.Right)) + 1

	return balance(n, key)
}

// Returns leftmost node in the subtree.
func leftmost(n *Node) *Node {
	for {
		if n.Left == nil {
			return n
		}
		n = n.Left
	}
	return n
}

// Returns node with the next key value or nil if the key is maximum.
func Next(n *Node, key int64) *Node {
	var prevRight *Node
	c := n
	for {
		if c.Key == key {
			if c.Right != nil {
				return leftmost(c.Right)
			}
			if prevRight != nil {
				return prevRight
			}
			break
		}
		if c.Key < key {
			c = c.Right
			continue
		}
		if c.Key > key {
			prevRight = c
			c = c.Left
			continue
		}
	}

	return nil
}
