package sort

import (
	"math/rand"
	"time"
)

const (
	DefaultCapacity = 10000  // 测试数组的长度
	DefaultRange    = 100000 // 数组元素大小范围
)

func GenerateRand(capacity int) []int {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	randomArr := make([]int, capacity)
	for i := 0; i < capacity; i++ {
		randomArr[i] = randSeed.Intn(DefaultRange)
	}
	return randomArr
}

func GetRandomIndex(capacity int) int {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	return randSeed.Intn(capacity)
}
