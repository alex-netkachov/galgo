package btree

// Binary tree node interface.
type IBTNode interface {
	// Returns -1 when node in parameter is greater, 0 if they are equal and 1
	// when the node is parameter is less then this node. In other words the sign
	// is equal to this node - compared node.
	Compare(IBTNode) int

	Left() IBTNode
	Right() IBTNode
	SetLeft(IBTNode)
	SetRight(IBTNode)
}

// Binary tree interface.
type IBTree interface {
	Root() IBTNode
	Insert(int)
}

// Simple implementation of the binary tree node with integer immutable keys
type BTNodeInt struct {
	key   int
	left  IBTNode
	right IBTNode
}

func CreateBTNodeInt(key int) *BTNodeInt {
	return &BTNodeInt{key: key}
}

func (this *BTNodeInt) Key() int {
	return this.key
}

func (this *BTNodeInt) Compare(node IBTNode) int {
	if n, ok := node.(*BTNodeInt); ok {
		if this.Key() < n.Key() {
			return -1
		}
		if this.Key() > n.Key() {
			return 1
		}
		return 0
	}
	panic("Expect non-null *BTNodeInt")
}

func (this *BTNodeInt) Left() IBTNode {
	return this.left
}

func (this *BTNodeInt) Right() IBTNode {
	return this.right
}

func (this *BTNodeInt) SetLeft(node IBTNode) {
	this.left = node
}

func (this *BTNodeInt) SetRight(node IBTNode) {
	this.right = node
}

// Simple implementation of the unbalanced binary tree.
type BTree struct {
	root IBTNode
}

func (this *BTree) Root() IBTNode {
	return this.root
}

func (this *BTree) Insert(node IBTNode) {
	if this.root == nil {
		this.root = node
		return
	}

	currentNode := this.root
	for {
		if currentNode.Compare(node) > 0 {
			if currentNode.Left() == nil {
				currentNode.SetLeft(node)
				break
			}
			currentNode = currentNode.Left()
		} else {
			if currentNode.Right() == nil {
				currentNode.SetRight(node)
				break
			}
			currentNode = currentNode.Right()
		}
	}
}
