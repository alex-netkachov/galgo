package btree

import (
	. "fmt"
	"testing"
)

func ExampleNext() {
	root := Insert(nil, 1)
	root = Insert(root, 2)
	root = Insert(root, 3)
	Println(Next(root, 2).Key)
	// 3
}

func TestNext(t *testing.T) {
	root := Insert(nil, 1)
	root = Insert(root, 2)
	root = Insert(root, 3)
	if Next(root, 1).Key != 2 {
		t.Error("Wrong next for 1")
	}
	if Next(root, 2).Key != 3 {
		t.Error("Wrong next for 2")
	}
	if Next(root, 3) != nil {
		t.Error("Wrong next for 3")
	}
}
