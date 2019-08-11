package gorm

import "testing"

func TestDelete(t *testing.T) {
	NewGorm()

	gormDB.Where("name=?", "fangfang").Delete(&User{})
}
