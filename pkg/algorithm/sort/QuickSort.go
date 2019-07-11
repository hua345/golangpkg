package sort

import "math/rand"

/**
快速排序算法首先会在序列中随机选择一个基准值（pivot），然后将除了基准值以外的数分
为“比基准值小的数”和“比基准值大的数”这两个类别，再将其排列成以下形式。
[ 比基准值小的数 ] 基准值 [ 比基准值大的数 ]
接着，对两个“[ ]”中的数据进行排序之后，整体的排序便完成了。对“[ ]”里面的数据
进行排序时同样也会使用快速排序。
*/
func QuickSort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	median := data[rand.Intn(len(data))]

	low_part := make([]int, 0, len(data))
	high_part := make([]int, 0, len(data))
	middle_part := make([]int, 0, len(data))

	for _, item := range data {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}

	low_part = QuickSort(low_part)
	high_part = QuickSort(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part
}
