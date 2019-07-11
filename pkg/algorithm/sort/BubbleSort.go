package sort

/**
冒泡排序就是重复“从序列右边开始比较相邻两个数字的大小，再根据结果交换两个数字
的位置”这一操作的算法。在这个过程中，数字会像泡泡一样，慢慢从右往左“浮”到序列的
顶端，所以这个算法才被称为“冒泡排序”。
*/
func BubbleSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
