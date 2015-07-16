package avltree

import (
	"testing"
)

func ExampleInsert() {
	root := Insert(nil, 1)
	root = Insert(root, 2)
	root = Insert(root, 3)
}

func TestInsert(t *testing.T) {
	root := Insert(nil, 1)
	root = Insert(root, 2)
	root = Insert(root, 3)
	if root.Key != 2 {
		t.Error("Tree is not balanced")
	}
}
