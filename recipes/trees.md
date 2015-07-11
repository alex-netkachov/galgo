Trees
===

General
---

### Struct

```go
type Node {
  Value int
  Nodes []*Node
}
```

Binary
---

### Struct

```go
type Node {
  Value int
  Left *Node
  Right *Node
}
```

AVL Tree
---

```go
// based on http://www.geeksforgeeks.org/avl-tree-set-1-insertion/

type Node struct {
	Value  int
	Left   *Node
	Right  *Node

	height int
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
	T2 := x.Right
	x.Right = y
	y.Left = T2
	y.height = max(height(y.Left), height(y.Right)) + 1
	x.height = max(height(x.Left), height(x.Right)) + 1
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
		return &Node{value, 1, nil, nil}
	}

	if value < n.Value {
		n.Left = insert(n.Left, value)
	} else {
		n.Right = insert(n.Right, value)
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
```
