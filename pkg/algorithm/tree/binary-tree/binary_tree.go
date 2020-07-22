package binarytree

import "fmt"

/**
二叉树(Binary tree)
二叉树是一种特殊的数据结构，最多有两个子节点：左子节点和右子节点

   1
 2   3
4 5   6
   7
    8
*/
type BNode struct {
	val   int
	left  *BNode
	right *BNode
}

type BTree struct {
	root *BNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newNode(val int) *BNode {
	n := &BNode{val, nil, nil}
	return n
}
func (node *BNode) Print() {
	fmt.Print(node.val, " ")
}

// 前序遍历： 根-> 左子树 -> 右子树
func PreOrder(node *BNode) {
	if node == nil {
		return
	}
	node.Print()
	PreOrder(node.left)
	PreOrder(node.right)
}

// 中序： 左子树-> 根 -> 右子树
func InOrder(node *BNode) {
	if node == nil {
		return
	}
	InOrder(node.left)
	node.Print()
	InOrder(node.right)
}

// 后序： 左子树-> 右子树 ->  根
func PostOrder(node *BNode) {
	if node == nil {
		return
	}
	PostOrder(node.left)
	PostOrder(node.right)
	node.Print()
}

//层次遍历(广度优先遍历)
func BreadthFirstSearch(root *BNode) {
	var q []*BNode // queue
	var n *BNode   // temporary BNode

	q = append(q, root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		n.Print()
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

// helper function for t.Depth
func _calculate_depth(bNode *BNode, depth int) int {
	if bNode == nil {
		return depth
	}
	return max(_calculate_depth(bNode.left, depth+1), _calculate_depth(bNode.right, depth+1))
}

func (btree *BTree) Depth() int {
	return _calculate_depth(btree.root, 0)
}
