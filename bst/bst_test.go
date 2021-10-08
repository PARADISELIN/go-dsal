package bst

import "testing"

func TestBSTInit(t *testing.T) {
	bst := BST{}

	if bst.root != nil {
		t.Error("when bst initialized, root node should be nil")
	}

	if bst.size != 0 {
		t.Error("when bst initialized, size should be zero")
	}
}

func TestBSTAddNonRecursive(t *testing.T) {
	bst := BST{}
	bst.AddNonRecursive(5)
	bst.AddNonRecursive(8)
	bst.AddNonRecursive(3)

	if bst.root.Val != 5 {
		t.Error("root value should be 5")
	}

	if bst.root.Right.Val != 8 {
		t.Error("right child value should be 8")
	}

	if bst.root.Left.Val != 3 {
		t.Error("left child value should be 3")
	}
}

func TestBSTAdd(t *testing.T) {
	bst := BST{}
	bst.Add(5)
	bst.Add(8)
	bst.Add(3)

	if bst.root.Val != 5 {
		t.Error("root value should be 5")
	}

	if bst.root.Right.Val != 8 {
		t.Error("right child value should be 8")
	}

	if bst.root.Left.Val != 3 {
		t.Error("left child value should be 3")
	}
}