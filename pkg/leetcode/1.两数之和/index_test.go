package main

import "testing"

func TestTwoSum(t *testing.T) {
	var nums = []int{3, 5, 9, 2, 7}
	t.Log(twoSum(nums, 9))
}

func TestTwoSumBest(t *testing.T) {
	var nums = []int{3, 5, 9, 2, 7}
	t.Log(twoSumBest(nums, 9))
}
