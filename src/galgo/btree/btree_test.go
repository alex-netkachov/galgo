package btree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tree := BTree{}
	tree.Insert(CreateBTNodeInt(1))
	tree.Insert(CreateBTNodeInt(2))
	tree.Insert(CreateBTNodeInt(3))

	if n, _ := tree.Root().(*BTNodeInt); n.Key() != 1 {
		t.Error("Wront root")
	}
	if tree.Root().Left() != nil {
		t.Error("Wront root left")
	}
	if tree.Root().Right() == nil {
		t.Error("Wront root right")
	}
}
