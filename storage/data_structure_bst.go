/*
	Using BST (Binary Search Tree) data structure, with average time complexity is O(log n) (https://en.wikipedia.org/wiki/Binary_search_tree)
*/

package storage

import (
	"errors"
	"fmt"
	"time"

	"github.com/Bhinneka/kece"
)

type node struct {
	left  *node
	right *node
	kece.Schema
}

// insert new node with recursive
func (n *node) insert(newNode *node) {
	if string(newNode.Key) > string(n.Key) {
		if n.right == nil {
			n.right = newNode
			return
		}
		n.right.insert(newNode)

	} else if string(newNode.Key) < string(n.Key) {
		if n.left == nil {
			n.left = newNode
			return
		}
		n.left.insert(newNode)

	} else {
		n.Value = newNode.Value
	}
}

// search node with key in tree nodes with recursive
func (n *node) searchNode(key string) *node {
	if n != nil {
		if key == string(n.Key) {
			return n
		} else if key > string(n.Key) {
			return n.right.searchNode(key)
		} else {
			return n.left.searchNode(key)
		}
	}
	return nil
}

// get lowest key in tree node (far left in nodes)
func (n *node) findLowestNode() *node {
	if n.left == nil {
		return n
	}
	return n.left.findLowestNode()
}

// get biggest key in tree node (far right in nodes)
func (n *node) findBiggestNode() *node {
	if n.right == nil {
		return n
	}
	return n.right.findBiggestNode()
}

func (n *node) replaceNode(parent, rep *node) {
	if n == parent.left {
		parent.left = rep
		return
	}
	parent.right = rep
}

func (n *node) delete(key string, parent *node) {
	if n == nil {
		return
	}

	switch {
	case key < string(n.Key):
		n.left.delete(key, n)
	case key > string(n.Key):
		n.right.delete(key, n)
	default:
		if n.left == nil && n.right == nil {
			n.replaceNode(parent, nil)
			return
		}

		if n.left == nil {
			n.replaceNode(parent, n.right)
			return
		}
		if n.right == nil {
			n.replaceNode(parent, n.left)
			return
		}
		replacement := n.left.findBiggestNode()
		n.Key = replacement.Key
		n.Value = replacement.Value
		n.Timestamp = replacement.Timestamp
		replacement.delete(string(replacement.Key), n)
	}
}

/*
TODO:
Want print tree:
        10
	  /    \
     8      14
    /  \
   7    9
*/

// printPreOrder print with traverse tree root->left->right
func (n *node) printPreOrder() {
	if n != nil {
		fmt.Println(string(n.Key), "=>", string(n.Value))
		n.left.printPreOrder()
		n.right.printPreOrder()
	}
}

/*
	------------------------------------
*/

// BST Binary Search Tree
type BST struct {
	root *node
}

// NewBST init new BST
func NewBST() kece.DataStructure {
	bst := new(BST)
	return bst
}

// Insert node with new key and value
func (tree *BST) Insert(key, value []byte) *kece.Schema {
	newSchema := &kece.Schema{Key: key, Value: value, Timestamp: time.Now()}
	newNode := new(node)
	newNode.Key = newSchema.Key
	newNode.Value = newSchema.Value
	newNode.Timestamp = newSchema.Timestamp
	if tree.root == nil {
		tree.root = newNode
		return newSchema
	}
	tree.root.insert(newNode)
	return newSchema
}

// Search node based on key
func (tree *BST) Search(key []byte) (*kece.Schema, error) {
	resNode := tree.root.searchNode(string(key))
	if resNode != nil {
		return &kece.Schema{Key: resNode.Key, Value: resNode.Value, Timestamp: resNode.Timestamp}, nil
	}
	return nil, errors.New(kece.ErrorEmptyValue)
}

// Delete node based on key
func (tree *BST) Delete(key []byte) error {
	if tree.root == nil {
		return errors.New(kece.ErrorEmptyValue)
	}
	tmpParent := &node{right: tree.root}
	tree.root.delete(string(key), tmpParent)
	if tmpParent.right == nil {
		tree.root = nil
	}
	return nil
}

// Print want show in pretty tree
func (tree *BST) Print() {
	tree.root.printPreOrder()
}
