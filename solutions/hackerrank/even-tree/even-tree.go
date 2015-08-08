package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go fmt even-tree.go && go run even-tree.go < even-tree.txt

type Node struct {
	ID    int
	Value int
}

type Edge struct {
	Node1 *Node
	Node2 *Node
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func (this *Graph) AddNode(id int, value int) {
	this.Nodes = append(this.Nodes, &Node{id, value})
}

func (this *Graph) RemoveNode(node *Node) {
	edges := this.GetIncidentEdges(node)
	for i := 0; i < len(this.Nodes); i++ {
		if this.Nodes[i] == node {
			this.Nodes = append(this.Nodes[:i], this.Nodes[i+1:]...)
		}
	}
	for _, e := range edges {
		this.RemoveEdge(e)
	}
}

func (this *Graph) IsLeaf(node *Node) bool {
	edges := this.GetIncidentEdges(node)
	return len(edges) == 1
}

func (this *Graph) GetNode(id int) *Node {
	for _, node := range this.Nodes {
		if node.ID == id {
			return node
		}
	}
	return nil
}

func (this *Graph) AddEdge(edge *Edge) {
	this.Edges = append(this.Edges, edge)
}

func (this *Graph) RemoveEdge(edge *Edge) {
	for i := 0; i < len(this.Edges); i++ {
		if this.Edges[i] == edge {
			this.Edges = append(this.Edges[:i], this.Edges[i+1:]...)
		}
	}
}

func (this *Graph) GetIncidentNodes(node *Node) []*Node {
	nodes := make([]*Node, 0)
	for _, edge := range this.Edges {
		if edge.Node1 == node {
			nodes = append(nodes, edge.Node2)
		} else if edge.Node2 == node {
			nodes = append(nodes, edge.Node1)
		}
	}
	return nodes
}

func (this *Graph) GetIncidentEdges(node *Node) []*Edge {
	edges := make([]*Edge, 0)
	for _, edge := range this.Edges {
		if edge.Node1 == node || edge.Node2 == node {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (this *Graph) getNodeIndex(node *Node) int {
	for i, n := range this.Nodes {
		if n == node {
			return i
		}
	}
	return -1
}

func (this *Graph) buildConnectedComponent(node *Node, index int, indices *[]int) {
	i := this.getNodeIndex(node)
	if (*indices)[i] == 0 {
		(*indices)[i] = index
		for _, n := range this.GetIncidentNodes(node) {
			this.buildConnectedComponent(n, index, indices)
		}
	}
}

func (this *Graph) GetConnectedComponents() [][]*Node {
	index := 0
	indices := make([]int, len(this.Nodes))
	for i := 0; i < len(this.Nodes); i++ {
		if indices[i] != 0 {
			continue
		}
		index++
		this.buildConnectedComponent(this.Nodes[i], index, &indices)
	}
	components := make([][]*Node, index)
	for i := 0; i < len(components); i++ {
		components[i] = make([]*Node, 0)
	}
	for i := 0; i < len(this.Nodes); i++ {
		components[indices[i]-1] = append(components[indices[i]-1], this.Nodes[i])
	}
	return components
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	s.Scan()
	m, _ := strconv.Atoi(s.Text())

	graph := Graph{}
	for i := 1; i <= n; i++ {
		graph.AddNode(i, 1)
	}

	for i := 0; i < m; i++ {
		s.Scan()
		id1, _ := strconv.Atoi(s.Text())
		s.Scan()
		id2, _ := strconv.Atoi(s.Text())

		graph.AddEdge(&Edge{graph.GetNode(id1), graph.GetNode(id2)})
	}

	removed := 0

	for {
		modified := false

		// 1. Collapse leafs with odd value
		nodes := make([]*Node, 0)
		nodes = append(nodes, graph.Nodes...)
		for _, n := range nodes {
			if graph.IsLeaf(n) && n.Value%2 == 1 {
				in := graph.GetIncidentNodes(n)[0]
				in.Value += n.Value
				graph.RemoveNode(n)
				modified = true
			}
		}

		// 2. Disconnect leafs with even values
		for _, n := range graph.Nodes {
			if graph.IsLeaf(n) && n.Value%2 == 0 {
				graph.RemoveEdge(graph.GetIncidentEdges(n)[0])
				modified = true
				removed++
			}
		}

		if !modified {
			break
		}
	}

	fmt.Printf("%v\n", removed)
}
