package postgres

import (
	"testing"
)

// row必须scan，不然会导致连接无法关闭，会一直占用连接，直到超过设置的生命周期
func TestSelectOne(t *testing.T) {
	NewPostgresDB()
	book := new(Book)
	row := postgresDB.QueryRow("select * from book where book_name=$1", "断舍离")
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	err := row.Scan(&book.Id, &book.BookName, &book.Price, &book.BookDesc)
	if err != nil {
		t.Error()
	}
	t.Log(book)
}
func TestSelectMulti(t *testing.T) {
	NewPostgresDB()
	book := new(Book)
	rows, err := postgresDB.Query("select * from book where book_name=$1", "断舍离")
	if err != nil {
		t.Error()
	}
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	index := 0
	for rows.Next() {
		//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
		index++
		if index >= 100 {
			break
		}
		err := rows.Scan(&book.Id, &book.BookName, &book.Price, &book.BookDesc)
		if err != nil {
			t.Error()
			return
		}
		t.Log(book)
	}
}

// 性能测试
//go test -bench=.
func BenchmarkSelect(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewPostgresDB()
	b.StartTimer()
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		book := new(Book)
		row := postgresDB.QueryRow("select * from book where book_name=$1", "断舍离")
		//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
		err := row.Scan(&book.Id, &book.BookName, &book.Price, &book.BookDesc)
		if err != nil {
			b.Error()
		}
	}
}

// 性能测试
//go test -bench=.
func BenchmarkSelectByIndex(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewPostgresDB()
	b.StartTimer()
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		book := new(Book)
		row := postgresDB.QueryRow("select * from book where id=$1", 38005)
		//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
		err := row.Scan(&book.Id, &book.BookName, &book.Price, &book.BookDesc)
		if err != nil {
			b.Error()
		}
	}
}

// 并发性能测试
//go test -bench=.
func BenchmarkSelectParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewPostgresDB()
	b.StartTimer()
	// 测试一个对象或者函数在多线程的场景下面是否安全
	book := new(Book)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			row := postgresDB.QueryRow("select * from book where book_name=$1", "断舍离")
			//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
			err := row.Scan(&book.Id, &book.BookName, &book.Price, &book.BookDesc)
			if err != nil {
				b.Error()
			}
		}
	})
}

// 并发性能测试
//go test -bench=.
func BenchmarkSelectByIndexParallel(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	NewPostgresDB()
	b.StartTimer()
	// 测试一个对象或者函数在多线程的场景下面是否安全
	book := new(Book)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			row := postgresDB.QueryRow("select * from book where id=$1", 38005)
			//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
			err := row.Scan(&book.Id, &book.BookName, &book.Price, &book.BookDesc)
			if err != nil {
				b.Error()
			}
		}
	})
}
