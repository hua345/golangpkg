package data_structures

import "testing"

func TestArray(t *testing.T) {
	// 数组默认初始化为0
	var initArr [3]int
	t.Log(initArr)
	// 数组字面值语法初始化数组
	var arr [3]int = [3]int{1, 2, 3}
	t.Log(arr)
	// 在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算。
	arr2 := [...]int{1, 2, 3}
	t.Logf("%T\n", arr2)
}
