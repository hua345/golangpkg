package sort

import "testing"

func TestGenerateRand(t *testing.T) {
	testData := GenerateRand(DefaultCapacity)
	if len(testData) != DefaultCapacity {
		t.Error("Generate Random Array Length Error")
	}
	for _, v := range testData {
		if v > DefaultRange {
			t.Error("Generate Random Number Range Error")
		}
	}
}
