package main

func twoSumBest(nums []int, target int) []int {
	elements := make(map[int]int)
	for k, v := range nums {
		if w, ok := elements[target-v]; ok {
			return []int{w, k}
		} else {
			elements[v] = k
		}
	}
	return []int{-1, -1}
}
