package postgres

import "testing"

func TestDelete(t *testing.T) {
	NewPostgresDB()
	// 预备表达式 用来优化SQL查询 提高性能 减少SQL注入的风险
	stmt, err := postgresDB.Prepare("delete from book where book_name=$1")
	if err != nil {
		t.Error(err)
	}
	result, err := stmt.Exec("断舍离")
	if err != nil {
		t.Error(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		t.Error(err)
	}
	t.Log("RowsAffected:", rowsAffected)
}
