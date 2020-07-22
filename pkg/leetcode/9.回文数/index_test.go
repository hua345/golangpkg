package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Run("测试121是回文数", func(t *testing.T) {
		isPalindromeStatus := isPalindrome(121)
		if !isPalindromeStatus {
			t.Error("121是回文数测试失败")
		}
	})
	t.Run("测试-121是回文数", func(t *testing.T) {
		isPalindromeStatus := isPalindrome(-121)
		if isPalindromeStatus {
			t.Error("-121是回文数测试失败")
		}
	})
	t.Run("测试10是回文数", func(t *testing.T) {
		isPalindromeStatus := isPalindrome(10)
		if isPalindromeStatus {
			t.Error("10是回文数测试失败")
		}
	})
}
