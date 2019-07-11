package sort

/**
归并排序算法会把序列分成长度相同的两个子序列，当无法继续往下分时（也就是每个子
序列中只有一个数据时），就对子序列进行归并。归并指的是把两个排好序的子序列合并成一个
有序序列。该操作会一直重复执行，直到所有子序列都归并为一个整体为止。
*/
func MergeSort(data []int) []int {

	if len(data) < 2 {
		return data
	}
	var middle = len(data) / 2
	var a = MergeSort(data[:middle])
	var b = MergeSort(data[middle:])
	return merge(a, b)
}

func merge(left []int, right []int) []int {

	var sortedData = make([]int, len(left)+len(right))
	var i = 0
	var j = 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sortedData[i+j] = left[i]
			i++
		} else {
			sortedData[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		sortedData[i+j] = left[i]
		i++
	}
	for j < len(right) {
		sortedData[i+j] = right[j]
		j++
	}
	return sortedData
}
