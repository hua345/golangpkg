package binary_search_tree

import (
	"fmt"
)

/**
二叉查找树(Binary Search Tree),又称二叉排序树(Binary Sort Tree)
若左子树非空，则左子树上的所有结点的关键字值均小于根结点的关键字值。
若右子树非空，则右子树上的所有结点的关键字值均大于根结点的关键字值。
左/右子树本身也分别是一棵二叉查找树

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

type BSTree struct {
	root *BNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newBNode(val int) *BNode {
	n := &BNode{val, nil, nil}
	return n
}
func (node *BNode) Print() {
	fmt.Print(node.val, " ")
}
func InsertNode(root *BNode, val int) *BNode {
	if root == nil {
		root = newBNode(val)
	}
	if val < root.val {
		root.left = InsertNode(root.left, val)
	} else {
		root.right = InsertNode(root.right, val)
	}
	return root
}
func inorderSuccessor(root *BNode) *BNode {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}
func DeleteNode(root *BNode, val int) *BNode {
	if root == nil {
		return nil
	}
	if val < root.val {
		root.left = DeleteNode(root.left, val)
	} else if val > root.val {
		root.right = DeleteNode(root.right, val)
	} else {
		// this is the node to delete
		// node with one child
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			n := root.right
			d := inorderSuccessor(n)
			d.left = root.left
			return root.right
		}
	}
	return root
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

func (btree *BSTree) Depth() int {
	return _calculate_depth(btree.root, 0)
}
