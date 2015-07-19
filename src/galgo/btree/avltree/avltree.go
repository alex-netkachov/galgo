package avltree

import (
	. "galgo/btree"
)

type IAVLBTNode interface {
	IBTNode
	Height() int
	SetHeight(int)
}

type AVLBTNodeInt struct {
	key    int
	left   IBTNode
	right  IBTNode
	height int
}

func CreateAVLBTNodeInt(key int) IBTNode {
	n := &AVLBTNodeInt{}
	n.key = key
	n.height = 1

	// initialise other properties
	// ...

	return n
}

func (this *AVLBTNodeInt) Key() int {
	return this.key
}

func (this *AVLBTNodeInt) Compare(node IBTNode) int {
	if n, ok := node.(*AVLBTNodeInt); ok {
		if this.Key() < n.Key() {
			return -1
		}
		if this.Key() > n.Key() {
			return 1
		}
		return 0
	}
	panic("Expect non-null *AVLBTNodeInt")
}

func (this *AVLBTNodeInt) Left() IBTNode {
	return this.left
}

func (this *AVLBTNodeInt) Right() IBTNode {
	return this.right
}

func (this *AVLBTNodeInt) SetLeft(node IBTNode) {
	this.left = node
}

func (this *AVLBTNodeInt) SetRight(node IBTNode) {
	this.right = node
}

func (this *AVLBTNodeInt) Height() int {
	return this.height
}

func (this *AVLBTNodeInt) SetHeight(height int) {
	this.height = height
}

type AVLBTree struct {
	root IBTNode
}

func (this *AVLBTree) Root() IBTNode {
	return this.root
}

func (this *AVLBTree) Insert(node IBTNode) {
	r, _ := this.root.(*AVLBTNodeInt)
	this.root = insert(r, node)
}

func (this *AVLBTree) InsertOrReplace(node IBTNode) {
	r, _ := this.root.(*AVLBTNodeInt)
	this.root = insertOrReplace(r, node)
}

func (this *AVLBTree) InsertOrIgnore(node IBTNode) {
	r, _ := this.root.(*AVLBTNodeInt)
	this.root = insertOrIgnore(r, node)
}

func height(node IBTNode) int {
	if node == nil {
		return 0
	}

	n, ok := node.(IAVLBTNode)
	if !ok {
		panic("Expect IAVLBTNode")
	}

	return n.Height()
}

func setHeight(node IBTNode, height int) {
	n, ok := node.(IAVLBTNode)
	if !ok {
		panic("Expect IAVLBTNode")
	}

	n.SetHeight(height)
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

func rotateLeft(x IBTNode) IBTNode {
	y := x.Right()
	t2 := y.Left()
	y.SetLeft(x)
	x.SetRight(t2)

	setHeight(x, max(height(x.Left()), height(x.Right()))+1)
	setHeight(y, max(height(y.Left()), height(y.Right()))+1)

	// update node invariants
	// ...

	return y
}

func rotateRight(y IBTNode) IBTNode {
	x := y.Left()
	t2 := x.Right()
	x.SetRight(y)
	y.SetLeft(t2)

	setHeight(y, max(height(y.Left()), height(y.Right()))+1)
	setHeight(x, max(height(x.Left()), height(x.Right()))+1)

	// update node invariants
	// ...

	return x
}

func balance(base IBTNode, node IBTNode) IBTNode {
	balance := height(base.Left()) - height(base.Right())

	if balance > 1 {
		switch node.Compare(base.Left()) {
		case -1:
			base.SetLeft(rotateLeft(base.Left()))
			return rotateRight(base)
		case 1:
			return rotateRight(base)
		}
	}

	if balance < -1 {
		switch node.Compare(base.Right()) {
		case -1:
			base.SetRight(rotateRight(base.Right()))
			return rotateLeft(base)
		case 1:
			return rotateLeft(base)
		}
	}

	return base
}

func insertOrIgnore(base IBTNode, node IBTNode) IBTNode {
	if base == nil || base == (*AVLBTNodeInt)(nil) {
		return node
	}

	// update statistics on the way down
	// ...

	switch base.Compare(node) {
	case -1:
		base.SetRight(insertOrIgnore(base.Right(), node))
	case 0:
		return base
	case 1:
		base.SetLeft(insertOrIgnore(base.Left(), node))
	}

	setHeight(base, max(height(base.Left()), height(base.Right()))+1)

	return balance(base, node)
}

func insertOrReplace(base IBTNode, node IBTNode) IBTNode {
	if base == nil || base == (*AVLBTNodeInt)(nil) {
		return node
	}

	// update statistics on the way down
	// ...

	switch base.Compare(node) {
	case -1:
		base.SetRight(insertOrReplace(base.Right(), node))
	case 0:
		setHeight(node, height(base))
		node.SetLeft(base.Left())
		node.SetRight(base.Right())

		// copy node statistics
		// ...

		return node
	case 1:
		base.SetLeft(insertOrReplace(base.Left(), node))
	}

	setHeight(base, max(height(base.Left()), height(base.Right()))+1)

	return balance(base, node)
}

func insert(base IBTNode, node IBTNode) IBTNode {
	if base == nil || base == (*AVLBTNodeInt)(nil) {
		return node
	}

	// update statistics on the way down
	// ...

	switch base.Compare(node) {
	case -1, 0:
		// duplicates go right
		base.SetRight(insert(base.Right(), node))
	case 1:
		base.SetLeft(insert(base.Left(), node))
	}

	setHeight(base, max(height(base.Left()), height(base.Right()))+1)

	return balance(base, node)
}
