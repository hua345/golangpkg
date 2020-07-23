package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2
	var result *ListNode
	if l1.Val > l2.Val {
		result = l2
	} else {
		result = l1
	}
	resultNode := result
	for node1 != nil {
		if node1.Next != nil {
			if node1.Next.Val > node2.Next.Val {
				resultNode.Next = node2.Next
				node2 = node2.Next
			} else {
				resultNode.Next = node1.Next
				node1 = node1.Next
			}
			resultNode = resultNode.Next
		} else {
			resultNode.Next = node2.Next
			node2 = node2.Next
			resultNode = resultNode.Next
			break
		}
	}
	if node2 != nil {
		resultNode.Next = node2
	}
	return resultNode
}
