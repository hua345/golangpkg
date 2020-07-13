package binarytree

import (
	"fmt"
	"testing"
)

/**
   1
 2   3
4 5   6
   7
    8
*/
func TestBTree(t *testing.T) {
	bTree := BTree{nil}
	bTree.root = newNode(1)
	bTree.root.left = newNode(2)
	bTree.root.right = newNode(3)
	bTree.root.left.left = newNode(4)
	bTree.root.left.right = newNode(5)
	bTree.root.right.right = newNode(6)
	bTree.root.right.right.right = newNode(7)
	bTree.root.right.right.right.right = newNode(8)

	fmt.Print("\n前序遍历: ")
	PreOrder(bTree.root)
	fmt.Print("\n中序遍历: ")
	InOrder(bTree.root)
	fmt.Print("\n后序遍历: ")
	PostOrder(bTree.root)
	fmt.Print("\n层次遍历: ")
	BreadthFirstSearch(bTree.root)
	t.Log(bTree.Depth())
}
