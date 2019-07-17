package redis

import (
	"sync"
)

// 参考: [Leaf——美团点评分布式ID生成系统](https://tech.meituan.com/2017/04/21/mt-leaf.html)
// Leaf这个名字是来自德国哲学家、数学家莱布尼茨的一句话：
// There are no two identical leaves in the world > “世界上没有两片相同的树叶”

type RedisLeaf struct {
	leafKey   string
	leafStart int64
	leafIndex int64
	queue     chan int64
	mu        *sync.Mutex
	running   bool
}

const defaultStep int64 = 1000

func NewLeaf(leafKey string) (leaf *RedisLeaf) {
	leafBegin, err := nextId(leafKey)
	if err != nil {
		panic(err)
	}
	leaf = &RedisLeaf{
		leafKey:   leafKey,
		leafStart: leafBegin * defaultStep,
		leafIndex: 0,
		mu:        new(sync.Mutex),
		running:   true,
		queue:     make(chan int64, 16),
	}
	go leaf.process()
	return
}

func (leaf *RedisLeaf) process() {
	defer func() { recover() }()
	for ; leaf.running; leaf.leafIndex++ {
		if leaf.leafIndex < defaultStep {
			leaf.queue <- leaf.leafStart + leaf.leafIndex
		} else {
			leaf.mu.Lock()
			leafBegin, err := nextId(leaf.leafKey)
			if err != nil {
				panic(err)
			}
			leaf.leafStart = leafBegin * defaultStep
			leaf.leafIndex = 0
			leaf.mu.Unlock()
		}
	}
}

func (leaf *RedisLeaf) NextId() int64 {
	return <-leaf.queue
}

func (leaf *RedisLeaf) Close() {
	leaf.running = false
	close(leaf.queue)
}

func nextId(idKey string) (int64, error) {
	orderId, err := RedisClient.IncrBy(idKey, 1).Result()
	if err != nil {
		return -1, err
	}
	return orderId, nil
}
