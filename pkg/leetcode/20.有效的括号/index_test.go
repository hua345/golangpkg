package main

import "testing"

func TestIsValid(t *testing.T) {
	t.Run("测试()有效的括号", func(t *testing.T) {
		isValidStatus := isValid("()")
		if !isValidStatus {
			t.Error("测试()有效的括号失败")
		}
	})
	t.Run("测试()[]{}有效的括号", func(t *testing.T) {
		isValidStatus := isValid("()[]{}")
		if !isValidStatus {
			t.Error("测试()[]{}有效的括号失败")
		}
	})
	t.Run("测试(]有效的括号", func(t *testing.T) {
		isValidStatus := isValid("(]")
		if isValidStatus {
			t.Error("测试(]有效的括号失败")
		}
	})
	t.Run("测试([)]有效的括号", func(t *testing.T) {
		isValidStatus := isValid("([)]")
		if isValidStatus {
			t.Error("测试([)]有效的括号失败")
		}
	})
	t.Run("测试{[]}有效的括号", func(t *testing.T) {
		isValidStatus := isValid("{[]}")
		if !isValidStatus {
			t.Error("测试{[]}有效的括号失败")
		}
	})
	t.Run("测试[有效的括号", func(t *testing.T) {
		isValidStatus := isValid("[")
		if isValidStatus {
			t.Error("测试[有效的括号失败")
		}
	})
	t.Run("测试]有效的括号", func(t *testing.T) {
		isValidStatus := isValid("]")
		if isValidStatus {
			t.Error("测试]有效的括号失败")
		}
	})
	t.Run("测试]()有效的括号", func(t *testing.T) {
		isValidStatus := isValid("]()")
		if isValidStatus {
			t.Error("测试]()有效的括号失败")
		}
	})
	t.Run("测试((有效的括号", func(t *testing.T) {
		isValidStatus := isValid("((")
		if isValidStatus {
			t.Error("测试((有效的括号失败")
		}
	})
	t.Run("测试((有效的括号", func(t *testing.T) {
		isValidStatus := isValid("{love[(fang)]}")
		if !isValidStatus {
			t.Error("测试((有效的括号失败")
		}
	})
}
