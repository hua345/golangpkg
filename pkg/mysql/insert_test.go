package mysql

import (
	"golangpkg/pkg/redis"
	"testing"
)

func TestInsert(t *testing.T) {
	NewMysqlDb()
	redis.NewRedis()
	leaf := redis.NewLeaf("bookKey")
	bookId := leaf.NextId()
	t.Log(bookId)
	result, err := mysqlDB.Exec("insert INTO book(id, book_name,price,book_desc) values(?,?,?,?)",
		bookId, "断舍离", "20.3", "断舍离是一本很好的书")
	if err != nil {
		t.Error(err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		t.Error(err)
	}
	t.Log("LastInsertID:", lastInsertID)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		t.Error(err)
	}
	t.Log("RowsAffected:", rowsAffected)
}

// 性能测试
//go test -bench=.
func BenchmarkInsert(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewMysqlDb()
	redis.NewRedis()
	leaf := redis.NewLeaf("bookKey")
	b.StartTimer()
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		bookId := leaf.NextId()
		result, err := mysqlDB.Exec("insert INTO book(id, book_name,price,book_desc) values(?,?,?,?)",
			bookId, "断舍离", "20.3", "断舍离是一本很好的书")
		if err != nil {
			b.Error(err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			b.Error(err)
		}
		if rowsAffected != int64(1) {
			b.Error("insert failed")
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkInsertParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewMysqlDb()
	redis.NewRedis()
	leaf := redis.NewLeaf("bookKey")
	b.StartTimer()
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bookId := leaf.NextId()
			result, err := mysqlDB.Exec("insert INTO book(id, book_name,price,book_desc) values(?,?,?,?)",
				bookId, "断舍离", "20.3", "断舍离是一本很好的书")
			if err != nil {
				b.Error(err)
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil {
				b.Error(err)
			}
			if rowsAffected != int64(1) {
				b.Error("insert failed")
			}
		}
	})
}
