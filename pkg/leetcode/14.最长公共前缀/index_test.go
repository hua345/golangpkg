package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("测试{'flower','flow','flight'}最长公共前缀", func(t *testing.T) {
		longestCommonPrefixStr := longestCommonPrefix([]string{"flower", "flow", "flight"})
		if longestCommonPrefixStr != "fl" {
			t.Error("测试{'flower','flow','flight'}最长公共前缀失败")
		}
	})
	t.Run("测试{'dog','racecar','car'}最长公共前缀", func(t *testing.T) {
		longestCommonPrefixStr := longestCommonPrefix([]string{"dog", "racecar", "car"})
		if longestCommonPrefixStr != "" {
			t.Error("测试{'dog','racecar','car'}最长公共前缀失败")
		}
	})
}
