package avltree

// go fmt . && go test . -v

// based on http://www.geeksforgeeks.org/avl-tree-set-1-insertion/

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	height int
}

func createNode(key int) *Node {
	n := &Node{}
	n.Key = key
	n.height = 1

	// initialise other properties
	// ...

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

/*
       y                           x
      / \     rotateRight()       / \
     x  T3    – – – – – – >     T1  y
    / \       < - - - - - -        / \
   T1  T2      rotateLeft()      T2  T3
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
	t2 := x.Right
	x.Right = y
	y.Left = t2
	y.height = max(height(y.Left), height(y.Right)) + 1
	x.height = max(height(x.Left), height(x.Right)) + 1

	// update node invariants
	// ...

	return x
}

func balance(n *Node, key int) *Node {
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

func InsertOrIgnore(n *Node, key int) *Node {
	if n == nil {
		return createNode(key)
	}

	// update statistics on the way down
	// ...

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

func InsertOrReplace(n *Node, key int) *Node {
	if n == nil {
		return createNode(key)
	}

	// update statistics on the way down
	// ...

	if key < n.Key {
		n.Left = InsertOrReplace(n.Left, key)
	} else if key > n.Key {
		n.Right = InsertOrReplace(n.Right, key)
	} else if key == n.Key {
		cn := createNode(key)
		cn.height = n.height
		cn.Left = n.Left
		cn.Right = n.Right

		// copy node statistics
		// ...

		return cn
	}

	n.height = max(height(n.Left), height(n.Right)) + 1

	return balance(n, key)
}

func Insert(n *Node, key int) *Node {
	if n == nil {
		return createNode(key)
	}

	// update statistics on the way down
	// ...

	if key < n.Key {
		n.Left = Insert(n.Left, key)
	} else {
		// duplicates go right
		n.Right = Insert(n.Right, key)
	}

	n.height = max(height(n.Left), height(n.Right)) + 1

	return balance(n, key)
}
