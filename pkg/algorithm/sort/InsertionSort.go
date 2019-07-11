package sort

/**
插入排序是一种从序列左端开始依次对数据进行排序的算法。在排序过程中，左侧的数据
陆续归位，而右侧留下的就是还未被排序的数据。插入排序的思路就是从右侧的未排序区域内
取出一个数据，然后将它插入到已排序区域内合适的位置上。
*/
func InsertionSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	for i := 1; i < len(data); i++ {
		for j := 0; j < i; j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
