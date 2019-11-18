package search

import (
	"github.com/hua345/golangpkg/pkg/algorithm/sort"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	testData := sort.GenerateRand(randomCapacity)
	randomIndex := sort.GetRandomIndex(randomCapacity)
	queryIndex := LinearSearch(testData, testData[randomIndex])
	if testData[randomIndex] != testData[queryIndex] {
		t.Log("randomValue: {} queryValue: {}", testData[randomIndex], testData[queryIndex])
		t.Error("Search Failed")
	}
}

//go test -bench=.
func BenchmarkLinearSearch(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	testData := sort.GenerateRand(randomCapacity)
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		randomIndex := sort.GetRandomIndex(randomCapacity)
		LinearSearch(testData, testData[randomIndex])
	}
}
