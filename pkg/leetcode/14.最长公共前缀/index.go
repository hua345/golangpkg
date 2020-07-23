package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefixStr string
	prefixStr = strs[0]
	for _, value := range strs[1:] {
		minLen := min(len(value), len(prefixStr))
		var prefixTemp []byte
		for i := 0; i < minLen; i++ {
			if value[i] == prefixStr[i] {
				prefixTemp = append(prefixTemp, value[i])
			} else {
				break
			}
		}
		prefixStr = string(prefixTemp)
	}
	return prefixStr
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
