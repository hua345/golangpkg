package main

import "testing"

func TestMergeTwoLists(t *testing.T) {
	t.Run("测试[1,2,4],[1,3,4]合并两个有序链表", func(t *testing.T) {
		l1 := &ListNode{1, nil}
		l1.Next = &ListNode{2, nil}
		l1.Next.Next = &ListNode{4, nil}
		l2 := &ListNode{1, nil}
		l2.Next = &ListNode{3, nil}
		l2.Next.Next = &ListNode{4, nil}
		temp := mergeTwoLists(l1, l2)
		for temp != nil && temp.Next != nil {
			if temp.Val > temp.Next.Val {
				t.Error("测试[1,2,4],[1,3,4]合并两个有序链表失败")
				break
			}
			temp = temp.Next
		}
	})
	t.Run("测试[],[]合并两个有序链表", func(t *testing.T) {
		var l1 *ListNode
		var l2 *ListNode
		temp := mergeTwoLists(l1, l2)
		for temp != nil && temp.Next != nil {
			if temp.Val > temp.Next.Val {
				t.Error("测试[],[]合并两个有序链表失败")
				break
			}
			temp = temp.Next
		}
	})
	t.Run("测试[2],[1]合并两个有序链表", func(t *testing.T) {
		l1 := &ListNode{2, nil}
		l2 := &ListNode{1, nil}
		temp := mergeTwoLists(l1, l2)
		for temp != nil && temp.Next != nil {
			if temp.Val > temp.Next.Val {
				t.Error("测试[2],[1]合并两个有序链表失败")
				break
			}
			temp = temp.Next
		}
	})
}
