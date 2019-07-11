package search

func BinarySearch(data []int, query int) int {
	return binarySearch(data, query, 0, len(data)-1)
}

func binarySearch(data []int, query int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	middle := int((lowIndex + highIndex) / 2)
	if data[middle] > query {
		return binarySearch(data, query, lowIndex, middle)
	} else if data[middle] < query {
		return binarySearch(data, query, middle+1, highIndex)
	} else {
		return middle
	}
}
