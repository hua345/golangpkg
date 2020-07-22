package main

import "testing"

func TestIsBalanced(t *testing.T) {
	bTreeTrue := &TreeNode{3, nil, nil}
	bTreeTrue.Left = &TreeNode{9, nil, nil}
	bTreeTrue.Right = &TreeNode{20, nil, nil}
	bTreeTrue.Right.Left = &TreeNode{15, nil, nil}
	bTreeTrue.Right.Right = &TreeNode{7, nil, nil}

	bTreeFalse := &TreeNode{1, nil, nil}
	bTreeFalse.Left = &TreeNode{2, nil, nil}
	bTreeFalse.Right = &TreeNode{2, nil, nil}
	bTreeFalse.Left.Left = &TreeNode{3, nil, nil}
	bTreeFalse.Left.Right = &TreeNode{3, nil, nil}
	bTreeFalse.Left.Left.Left = &TreeNode{4, nil, nil}
	bTreeFalse.Left.Left.Right = &TreeNode{4, nil, nil}

	t.Run("测试正确平衡二叉树", func(t *testing.T) {
		status := isBalanced(bTreeTrue)
		if !status {
			t.Error("平衡二叉树检查失败")
		}
	})
	t.Run("测错误平衡二叉树", func(t *testing.T) {
		status := isBalanced(bTreeFalse)
		if status {
			t.Error("平衡二叉树检查失败")
		}
	})
}
