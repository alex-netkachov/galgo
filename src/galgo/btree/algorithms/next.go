package algorithms

import (
	. "galgo/btree"
)

// Returns leftmost node in the subtree.
func Leftmost(n IBTNode) IBTNode {
	for {
		if n.Left() == nil {
			return n
		}
		n = n.Left()
	}
	return n
}

// Returns rightmost node in the subtree.
func Rightmost(n IBTNode) IBTNode {
	for {
		if n.Right() == nil {
			return n
		}
		n = n.Right()
	}
	return n
}

// Returns node with the next key value or nil if the key is maximum within
// the subtree.
func Next(root IBTNode, node IBTNode) IBTNode {
	var prevRight IBTNode

	current := root
	for current != nil {
		switch current.Compare(node) {
		case -1:
			current = current.Right()
		case 0:
			if current.Right() != nil {
				return Leftmost(current.Right())
			}
			if prevRight != nil {
				return prevRight
			}
			return nil
		case 1:
			prevRight = current
			current = current.Left()
		}
	}

	return nil
}
