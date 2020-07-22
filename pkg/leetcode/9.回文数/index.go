package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	return isPalindrome2(x)
}
func isPalindrome1(x int) bool {
	if x >= 0 && x < 10 {
		return true
	}
	if x < 0 {
		return false
	}
	var data []byte = []byte(strconv.Itoa(x))
	var status = true
	middleLen := len(data) / 2
	for i := 0; i < middleLen; i++ {
		if data[i] != data[len(data)-i-1] {
			status = false
			break
		}
	}
	return status
}
func isPalindrome2(x int) bool {
	if x >= 0 && x < 10 {
		return true
	}
	if x < 0 {
		return false
	}
	ori := x
	var n int
	for x != 0 {
		n = n*10 + x%10
		x = x / 10
	}
	return ori == n
}
