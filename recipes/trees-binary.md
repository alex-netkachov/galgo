Model
---

Basic model for binary search tree.

```go
type Node {
  // main property
  Value int

  // navigation
  Left *Node
  Right *Node

  // other properties
  // ...
}
```

AVL Tree
---

Balanced binary search tree.

To change the type of the main property replace `Value int` and `func Insert(n *Node, value int) *Node` with `Value type` and `func Insert(n *Node, value type) *Node`. The `type` should support the operation `<`.


```go
// based on http://www.geeksforgeeks.org/avl-tree-set-1-insertion/

type Node struct {
	// main property
	Value int

	// navigation
	Left  *Node
	Right *Node

	// other properties
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
	
	// update node invariants
	// ...

	return y
}

func rotateRight(y *Node) *Node {
	x := y.Left
	T2 := x.Right
	x.Right = y
	y.Left = T2
	y.height = max(height(y.Left), height(y.Right)) + 1
	x.height = max(height(x.Left), height(x.Right)) + 1
	
	// update node invariants
	// ...

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
		return &Node{value, nil, nil, 1}
	}

	if value < n.Value {
		n.Left = Insert(n.Left, value)
	} else {
		n.Right = Insert(n.Right, value)
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

Example:

```go
r := tree.Insert(nil, 1)
r = tree.Insert(r, 2)
r = tree.Insert(r, 3)
```

### Invariants

#### Count of Subnodes

```go
type Node struct {
	...

	count  int
}

func count(n *Node) int {
	if n == nil {
		return 0
	}
	return n.count
}

func rotateLeft(x *Node) *Node {
	...

	cx := count(x) - count(x.Left) - count(y)
	cy := count(y) - count(x.Right) - count(y.Right)

	x.count = cx + count(x.Left) + count(x.Right)
	y.count = cy + count(y.Left) + count(y.Right)

	return y
}

func rotateRight(y *Node) *Node {
	...

	cx := count(x) - count(x.Left) - count(y.Left)
	cy := count(y) - count(x) - count(y.Right)
	y.count = cy + count(y.Left) + count(y.Right)
	x.count = cx + count(x.Left) + count(x.Right)

	return x
}

func Insert(n *Node, value int) *Node {
	// 1. Perform the normal BST rotation
	if n == nil {
		return &Node{value, nil, nil, 1, 1}
	}

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
```
