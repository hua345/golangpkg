package main

import "testing"

func TestReverse(t *testing.T) {
	t.Run("测试123整数反转", func(t *testing.T) {
		reverseNum := reverse(123)
		if reverseNum != 321 {
			t.Error("123整数反转测试失败")
		}
	})
	t.Run("测试以120整数反转", func(t *testing.T) {
		reverseNum := reverse(120)
		if reverseNum != 21 {
			t.Error("120整数反转测试失败")
		}
	})
	t.Run("测试901000整数反转", func(t *testing.T) {
		reverseNum := reverse(901000)
		if reverseNum != 109 {
			t.Error("901000整数反转测试失败")
		}
	})
	t.Run("测试1534236469整数反转", func(t *testing.T) {
		reverseNum := reverse(1534236469)
		if reverseNum != 0 {
			t.Error("1534236469整数反转测试失败")
		}
	})

	t.Run("测试-123整数反转", func(t *testing.T) {
		reverseNum := reverse(-123)
		if reverseNum != -321 {
			t.Error("-123整数反转测试失败")
		}
	})

}
