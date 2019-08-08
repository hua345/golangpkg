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
	result, err := mysqlDB.Exec("insert INTO book(id, book_name,price,book_desc) values(?,?,?,?)", bookId, "断舍离", "20.3", "断舍离是一本很好的书")
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
