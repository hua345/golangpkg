package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Len := 1
	nextNode := l1
	for nextNode.Next != nil {
		nextNode = nextNode.Next
		l1Len++
	}
	l2Len := 1
	nextNode = l2
	for nextNode.Next != nil {
		nextNode = nextNode.Next
		l2Len++
	}
	resultNode := &ListNode{Val: 0}
	if l1Len >= l2Len {
		addNode(l1, l2, resultNode)
	} else {
		addNode(l2, l1, resultNode)
	}

	return resultNode
}

func addNode(longNode *ListNode, shortNode *ListNode, resultNode *ListNode) {
	nowNode := resultNode
	for ; ; longNode = longNode.Next {
		if shortNode != nil {
			num := nowNode.Val + longNode.Val + shortNode.Val
			if num >= 10 {
				nowNode.Val = num % 10
				nowNode.Next = &ListNode{Val: 1}
				nowNode = nowNode.Next
			} else {
				if longNode.Next != nil {
					nowNode.Val = num
					nowNode.Next = &ListNode{Val: 0}
					nowNode = nowNode.Next
				} else {
					nowNode.Val = num
				}
			}
			shortNode = shortNode.Next
		} else {
			num := nowNode.Val + longNode.Val
			if num >= 10 {
				nowNode.Val = num % 10
				nowNode.Next = &ListNode{Val: 1}
				nowNode = nowNode.Next
			} else {
				if longNode.Next != nil {
					nowNode.Val = num
					nowNode.Next = &ListNode{Val: 0}
					nowNode = nowNode.Next
				} else {
					nowNode.Val = num
				}
			}
		}
		if longNode.Next == nil {
			break
		}
	}
}
