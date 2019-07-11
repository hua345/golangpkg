package sort

/**
选择排序就是重复“从待排序的数据中寻找最小值，将其与序列最左边的数字进行交换”
这一操作的算法。在序列中寻找最小值时使用的是线性查找。
*/
func SelectionSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	for i := 0; i < len(data)-1; i++ {
		minIndex := 0
		for j := i + 1; j < len(data); j++ {
			if data[minIndex] > data[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
	return data
}
