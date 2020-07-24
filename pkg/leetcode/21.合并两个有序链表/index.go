package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoLists2(l1, l2)
}
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	node1 := l1
	node2 := l2
	var result *ListNode
	if l1.Val <= l2.Val {
		result = &ListNode{l1.Val, nil}
		node1 = l1.Next
	} else {
		result = &ListNode{l2.Val, nil}
		node2 = l2.Next
	}
	resultNode := result
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			resultNode.Next = &ListNode{node1.Val, nil}
			node1 = node1.Next
		} else {
			resultNode.Next = &ListNode{node2.Val, nil}
			node2 = node2.Next
		}
		resultNode = resultNode.Next
	}
	if node1 != nil {
		resultNode.Next = node1
	}
	if node2 != nil {
		resultNode.Next = node2
	}
	return result
}
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummyHead := new(ListNode)
	head := dummyHead
	dummyHead.Next = head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			head.Next = l1
			head = head.Next
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
			head = head.Next
		}
	}
	if l1 == nil {
		head.Next = l2
	}
	if l2 == nil {
		head.Next = l1
	}
	return dummyHead.Next
}
