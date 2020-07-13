package data_structures

import "testing"

/**
Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。
一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已。
一个slice由三个部分构成：指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址，要注意的是slice的第一个元素并不一定就是数组的第一个元素。
多个slice之间可以共享底层的数据，并且引用的数组部分区间可能重叠。
*/
func TestSlice(t *testing.T) {
	months := [...]string{1: "January", 2: "FEBRUARY", 3: "MARCH", 4: "APRIL", 5: "MAY", 6: "JUNE", 7: "JULY", 8: "AUGUST", 9: "SEPTEMBER", 10: "OCTOBER", 11: "NOVEMBER", 12: "December"}

	Q2 := months[4:7]
	summer := months[6:9]
	t.Log(Q2)
	t.Log(summer)
}

/**
内置的append函数用于向slice追加元素
每次调用append函数，必须先检测slice底层数组是否有足够的容量来保存新添加的元素
每一次容量的变化都会导致重新分配内存和copy操作：
*/
func TestSliceAppend(t *testing.T) {
	var mySlice []int
	for i := 0; i < 10; i++ {
		mySlice = append(mySlice, i)
		t.Logf("%d cap=%d\t%v\n", len(mySlice), cap(mySlice), mySlice)
	}
}
