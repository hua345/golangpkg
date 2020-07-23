package main

func romanToInt(s string) int {
	return romanToInt2(s)
}
func romanToInt1(s string) int {
	var resultNum int = 0
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var data []byte = []byte(s)
	if len(data) == 1 {
		value, ok := romanMap[data[0]]
		if ok {
			resultNum = value
		}
		return resultNum
	}
	for i := len(data) - 1; i >= 1; i-- {
		nextValue, ok := romanMap[data[i]]
		preValue, preOk := romanMap[data[i-1]]
		if ok && preOk {
			if preValue < nextValue {
				resultNum = resultNum + (nextValue - preValue)
				i--
				if 1 == i {
					value, ok := romanMap[data[0]]
					if ok {
						resultNum = resultNum + value
					}
				}
			} else {
				if 0 == i-1 {
					resultNum = resultNum + preValue + nextValue
				} else {
					resultNum = resultNum + nextValue
				}
			}
		}
	}
	return resultNum
}
func romanToInt2(s string) int {
	var resultNum int = 0
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := romanMap[s[i]]
		if cur >= pre {
			resultNum += cur
		} else {
			resultNum -= cur
		}
		pre = cur
	}
	return resultNum
}
