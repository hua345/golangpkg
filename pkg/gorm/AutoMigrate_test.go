package gorm

import "testing"

func TestAutoMigrate(t *testing.T) {
	GetInstance().AutoMigrate(&User{})
	GetInstance().AutoMigrate(&Book{})
}
