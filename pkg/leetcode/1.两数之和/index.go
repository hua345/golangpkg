package main

func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	for index := 0; index < len(nums); index++ {
		for nextIndex := index + 1; nextIndex < len(nums); nextIndex++ {
			if nums[index]+nums[nextIndex] == target {
				return []int{index, nextIndex}
			}
		}
	}
	return []int{}
}
