package mysql

import "testing"

func TestDelete(t *testing.T) {
	NewMysqlDb()
	stmt, err := mysqlDB.Prepare("delete from book where book_name=?")
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
