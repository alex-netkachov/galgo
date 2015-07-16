Model
---

Basic model for binary search tree.

```go
type Node {
	// main property
	Key    int

	Left *Node
	Right *Node

	// other properties
	// ...
}
```

AVL Tree
---

Balanced binary search tree.

To change the type of the main property replace `Key int` and `func Insert(n *Node, key int) *Node` with `Key type` and `func Insert(n *Node, key type) *Node`. The `type` should support the operations `<`, `>`, and `==`.

See the file [avltree/avltree.go](.../avltree/avltree.go) for implementation of the AVL tree in Go.

Example:

```go
r := avltree.Insert(nil, 1)
r = avltree.Insert(r, 2)
r = avltree.Insert(r, 3)
```

### Count of Subnodes

This is a special version of the AVL tree that counts the number of subnodes for every node.

```go
type Node struct {
	...
	Count  int
}

func createNode(key int) *Node {
	...
	n.Count = 1
	return n
}

func count(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Count
}

func rotateLeft(x *Node) *Node {
	...

	cx := count(x) - count(x.Left) - count(y)
	cy := count(y) - count(x.Right) - count(y.Right)

	x.Count = cx + count(x.Left) + count(x.Right)
	y.Count = cy + count(y.Left) + count(y.Right)

	return y
}

func rotateRight(y *Node) *Node {
	...

	cx := count(x) - count(x.Left) - count(y.Left)
	cy := count(y) - count(x) - count(y.Right)
	y.Count = cy + count(y.Left) + count(y.Right)
	x.Count = cx + count(x.Left) + count(x.Right)

	return x
}

func InsertXXX(n *Node, value int) *Node {
	if n == nil {
		return &Node{value, nil, nil, 1, 1}
	}

	// update statistics
	n.Count++

	...
}
```

`rank` calculates order of the node in the tree:

```go
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
			s += c.Count - count(c.Right)
			c = c.Right
			continue
		}
		if v == c.Value {
			s += c.Count - count(c.Right)
			return s
		}
	}
	return -1
}
```
