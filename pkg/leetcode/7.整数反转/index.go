package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	return reverse2(x)
}
func reverse1(x int) int {
	if 0 == x {
		return 0
	}
	str := strconv.Itoa(x)
	var data []byte = []byte(str)
	var result []byte
	if '-' == data[0] {
		result = append(result, '-')
		data = data[1:]
	}
	var lastZero bool = true
	for i := len(data) - 1; i >= 0; i-- {
		if lastZero && '0' == data[i] {
			continue
		}
		lastZero = false
		result = append(result, data[i])
	}
	dataStr := string(result)
	dataNum, err := strconv.Atoi(dataStr)
	if err != nil {
		return 0
	}
	if dataNum > math.MaxInt32 || dataNum < math.MinInt32 {
		return 0
	}
	return dataNum
}
func reverse2(x int) int {
	var dataNum int64
	for x != 0 {
		dataNum = dataNum*10 + int64(x%10)
		x = x / 10
	}
	if dataNum > math.MaxInt32 || dataNum < math.MinInt32 {
		return 0
	}
	return int(dataNum)
}
