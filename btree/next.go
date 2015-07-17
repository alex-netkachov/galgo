package btree

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
func Next(n *Node, key int) *Node {
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
