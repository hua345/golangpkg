package main

import "testing"

func TestRomanToInt(t *testing.T) {
	t.Run("测试罗马数III", func(t *testing.T) {
		romanInt := romanToInt("III")
		if romanInt != 3 {
			t.Errorf("罗马数III得到%d测试失败", romanInt)
		}
	})
	t.Run("测试罗马数IV", func(t *testing.T) {
		romanInt := romanToInt("IV")
		if romanInt != 4 {
			t.Errorf("罗马数IV得到%d测试失败", romanInt)
		}
	})
	t.Run("测试罗马数IX", func(t *testing.T) {
		romanInt := romanToInt("IX")
		if romanInt != 9 {
			t.Errorf("罗马数IX得到%d测试失败", romanInt)
		}
	})
	t.Run("测试罗马数MCMXCIV", func(t *testing.T) {
		romanInt := romanToInt("MCMXCIV")
		if romanInt != 1994 {
			t.Errorf("罗马数MCMXCIV得到%d测试失败", romanInt)
		}
	})
}
