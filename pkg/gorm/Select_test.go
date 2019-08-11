package gorm

import "testing"

func TestSelect(t *testing.T) {
	NewGorm()

	var user User
	gormDB.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	t.Log(user)

	// 通过主键查询最后一条记录
	user = User{}
	gormDB.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1
	t.Log(user)

	// Get first matched record
	user = User{}
	gormDB.Where("name = $1", "fangfang").First(&user)
	// SELECT * FROM users WHERE name = 'fangfang' limit 1;
	t.Log(user)
	// LIKE
	user = User{}
	gormDB.Where("name LIKE ?", "%fang%").First(&user)
	t.Log(user)
	// AND
	user = User{}
	result := gormDB.Where("name = ? AND age >= ?", "fangfang", "22").Find(&user)
	t.Log("RowsAffected", result.RowsAffected)
	t.Log(user)
	// Struct
	user = User{}
	result = gormDB.Where(&User{Name: "fang", Age: 24}).First(&user)
	t.Log("RowsAffected", result.RowsAffected)
	// SELECT * FROM users WHERE name = "fang" AND age = 24 LIMIT 1;
	t.Log(user)
}

// 性能测试
//go test -bench=.
func BenchmarkSelect(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewGorm()
	b.StartTimer()
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		var user User
		err := gormDB.Where("name = $1", "fangfang").First(&user).Error
		if err != nil {
			b.Error(err)
		}
	}
}

// 性能测试
//go test -bench=.
func BenchmarkSelectByIndex(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewGorm()
	b.StartTimer()
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		var user User
		err := gormDB.Where("id = ?", 2020).First(&user).Error
		if err != nil {
			b.Error(err)
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkSelectParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewGorm()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var user User
			err := gormDB.Where("name = $1", "fangfang").First(&user).Error
			if err != nil {
				b.Error(err)
			}
		}
	})
}

// 并发性能测试
//go test -bench=.
func BenchmarkSelectByIndexParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewGorm()
	b.StartTimer()
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var user User
			err := gormDB.Where("id = ?", 2020).First(&user).Error
			if err != nil {
				b.Error(err)
			}
		}
	})
}
