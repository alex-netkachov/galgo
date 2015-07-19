package avltree

import (
	"testing"
)

func ExampleInsert() {
	tree := AVLBTree{}
	tree.Insert(CreateAVLBTNodeInt(1))
	tree.Insert(CreateAVLBTNodeInt(2))
	tree.Insert(CreateAVLBTNodeInt(3))
}

func TestInsert(t *testing.T) {
	tree := &AVLBTree{}
	tree.Insert(CreateAVLBTNodeInt(1))
	tree.Insert(CreateAVLBTNodeInt(2))
	tree.Insert(CreateAVLBTNodeInt(3))

	if n, _ := tree.Root().(*AVLBTNodeInt); n.Key() != 2 {
		t.Error("Tree is not balanced")
	}
}
