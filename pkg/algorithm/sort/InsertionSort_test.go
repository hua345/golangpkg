package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	testData := GenerateRand(DefaultCapacity)
	sortData := InsertionSort(testData)
	bubbleSortData := BubbleSort(testData)
	for i, v := range sortData {
		if bubbleSortData[i] != v {
			t.Error("InsertionSort Result Not Equal BubbleSort Result")
		}
	}
}

//go test -bench=.
func BenchmarkInsertionSort(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	testData := GenerateRand(DefaultCapacity)
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		InsertionSort(testData)
	}
}
