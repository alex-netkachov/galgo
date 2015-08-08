package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

// go fmt cut-the-tree.go && go run cut-the-tree.go < cut-the-tree.txt

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	keys := ReadInts(s)
	nodes := make([]INode, len(keys))
	for i := 0; i < len(keys); i++ {
		nodes[i] = INode(&NodeInt{i + 1, keys[i], []INode{}, 0})
	}

	for i := 0; i < len(nodes)-1; i++ {
		n1 := ReadInt(s) - 1
		n2 := ReadInt(s) - 1
		nodes[n1].Append(nodes[n2])
		nodes[n2].Append(nodes[n1])
	}

	vmask := make([]bool, len(nodes))
	vlist := make([]INode, len(nodes))
	vi := 0
	vlist[vi] = nodes[0]
	vmask[vi] = true

	for i := 0; i < len(nodes); i++ {
		n := vlist[i]
		j := 0
		for j < len(n.Nodes()) {
			sn, _ := n.Nodes()[j].(*NodeInt)
			if vmask[sn.ID()-1] {
				vlist[i].Remove(sn)
			} else {
				vi++
				vlist[vi] = sn
				vmask[sn.ID()-1] = true
				j++
			}
		}
	}

	root, _ := nodes[0].(*NodeInt)

	DFS(root, func(node INode) {
		n, _ := node.(*NodeInt)
		n.sum = n.key
		for _, sn := range node.Nodes() {
			sn1, _ := sn.(*NodeInt)
			n.sum += sn1.sum
		}
	})

	min := root.sum

	DFS(root, func(node INode) {
		n, _ := node.(*NodeInt)
		d := root.sum - n.sum - n.sum
		if d < 0 {
			d = -d
		}
		if d < min {
			min = d
		}
	})

	Println(min)
}

func ReadInt(s *bufio.Scanner) int {
	s.Scan()
	i, _ := strconv.Atoi(s.Text())
	return i
}

func ReadInts(s *bufio.Scanner) []int {
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}
	return a
}

func DFS(node INode, callback func(INode)) {
	for _, n := range node.Nodes() {
		DFS(n, callback)
	}
	callback(node)
}

type INode interface {
	Nodes() []INode
	Append(INode)
	Remove(INode)
}

type NodeInt struct {
	id    int
	key   int
	nodes []INode
	sum   int
}

func (this *NodeInt) ID() int {
	return this.id
}

func (this *NodeInt) Key() int {
	return this.key
}

func (this *NodeInt) Nodes() []INode {
	return this.nodes
}

func (this *NodeInt) Append(node INode) {
	this.nodes = append(this.nodes, node)
}

func (this *NodeInt) Remove(node INode) {
	n, _ := node.(*NodeInt)
	for k, v := range this.nodes {
		if v == n {
			this.nodes = append(this.nodes[0:k], this.nodes[k+1:]...)
			return
		}
	}
}
