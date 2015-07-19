package algorithms

import (
	. "fmt"
	. "galgo/btree"
	"testing"
)

func ExampleNext() {
	tree := BTree{}

	tree.Insert(CreateBTNodeInt(1))

	n2 := CreateBTNodeInt(2)
	tree.Insert(n2)

	tree.Insert(CreateBTNodeInt(3))

	node, _ := Next(tree.Root(), n2).(*BTNodeInt)
	Println(node.Key())
	// 3
}

func TestNext(t *testing.T) {
	tree := BTree{}

	n1 := CreateBTNodeInt(1)
	tree.Insert(n1)

	n2 := CreateBTNodeInt(2)
	tree.Insert(n2)

	n3 := CreateBTNodeInt(3)
	tree.Insert(n3)

	if n, _ := Next(tree.Root(), n1).(*BTNodeInt); n.Key() != 2 {
		t.Error("Wrong next for 1")
	}
	if n, _ := Next(tree.Root(), n2).(*BTNodeInt); n.Key() != 3 {
		t.Error("Wrong next for 2")
	}
	if Next(tree.Root(), n3) != nil {
		t.Error("Wrong next for 3")
	}
}
