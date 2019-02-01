//credit Agung Dwi Prasetyo https://github.com/agungdwiprasetyo
package storage

import (
	"testing"
)

func TestNewBST(t *testing.T) {
	bst := NewBST()
	if bst == nil {
		t.Errorf("bst is nil")
	}
}

func TestBST_Insert(t *testing.T) {
	bst := new(BST)
	bst.Insert([]byte("d"), []byte("ini d"))
	bst.Insert([]byte("b"), []byte("ini b"))
	bst.Insert([]byte("c"), []byte("ini c"))
	bst.Insert([]byte("e"), []byte("ini e"))
	bst.Insert([]byte("a"), []byte("ini a"))

	root := bst.root
	if string(root.Key) != "d" {
		t.Errorf("should d")
	}
	leftNode := root.left
	if string(leftNode.Key) != "b" {
		t.Errorf("should b")
	}
	rightNode := root.right
	if string(rightNode.Key) != "e" {
		t.Errorf("should e")
	}
}

func TestBST_Search(t *testing.T) {
	bst := new(BST)
	bst.Insert([]byte("d"), []byte("ini d"))
	bst.Insert([]byte("b"), []byte("ini b"))
	bst.Insert([]byte("c"), []byte("ini c"))
	bst.Insert([]byte("e"), []byte("ini e"))
	bst.Insert([]byte("a"), []byte("ini a"))

	res, err := bst.Search([]byte("e"))
	if err != nil {
		t.Errorf("should no error")
	}
	if string(res.Value) != "ini e" {
		t.Errorf("should 'ini e'")
	}
}

func TestBST_Delete(t *testing.T) {
	bst := new(BST)
	bst.Insert([]byte("d"), []byte("ini d"))
	bst.Insert([]byte("b"), []byte("ini b"))
	bst.Insert([]byte("c"), []byte("ini c"))
	bst.Insert([]byte("e"), []byte("ini e"))
	bst.Insert([]byte("a"), []byte("ini a"))

	err := bst.Delete([]byte("e"))
	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}
	rightRoot := bst.root.right
	if rightRoot != nil {
		t.Errorf("should nil")
	}

	err = bst.Delete([]byte("b"))
	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}
	leftRoot := bst.root.left
	if string(leftRoot.Key) != "a" {
		t.Errorf("expected a, result is %s", string(leftRoot.Key))
	}
}
