package main

import "testing"

func displayNode(result *ListNode, t *testing.T) {
	for ; ; result = result.Next {
		t.Log(result.Val)
		if result.Next == nil {
			break
		}
	}
}
func TestAddTwoNumbers(t *testing.T) {
	aa := &ListNode{2, &ListNode{Val: 3}}
	bb := &ListNode{9, &ListNode{1, &ListNode{2, &ListNode{Val: 2}}}}
	result := addTwoNumbers(aa, bb)
	displayNode(bb, t)
	displayNode(result, t)
}
