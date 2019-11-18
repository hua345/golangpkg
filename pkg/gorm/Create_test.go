package gorm

import (
	"github.com/hua345/golangpkg/pkg/redis"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	NewGorm()
	redis.NewRedis()
	leaf := redis.NewLeaf("bookKey")
	bookId := leaf.NextId()

	user := &User{
		Name:       "fangfang",
		UserMobile: "17875266970",
		Age:        24,
	}
	user.ID = bookId
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	t.Log(user.CreatedTime.Format("2006/01/02 15:04:05"))
	err := gormDB.Create(&user).Error
	if err != nil {
		t.Error(err)
	}
}

// 性能测试
//go test -bench=.
func BenchmarkInsert(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewGorm()
	redis.NewRedis()
	leaf := redis.NewLeaf("userKey")
	b.StartTimer()
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		user := &User{
			Name:       "fangfang",
			UserMobile: "17875266970",
			Age:        24,
		}
		user.CreatedTime = time.Now()
		user.UpdatedTime = time.Now()
		bookId := leaf.NextId()
		user.ID = bookId
		err := gormDB.Create(&user).Error
		if err != nil {
			b.Error(err)
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkInsertParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewGorm()
	redis.NewRedis()
	leaf := redis.NewLeaf("userKey")
	b.StartTimer()
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			user := &User{
				Name:       "fangfang",
				UserMobile: "17875266970",
				Age:        24,
			}
			user.CreatedTime = time.Now()
			user.UpdatedTime = time.Now()
			bookId := leaf.NextId()
			user.ID = bookId
			err := gormDB.Create(&user).Error
			if err != nil {
				b.Error(err)
			}
		}
	})
}
