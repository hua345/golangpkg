package gorm

import "testing"

func TestAutoMigrate(t *testing.T) {
	NewGorm()
	gormDB.AutoMigrate(&User{})
}
