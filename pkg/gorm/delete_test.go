package gorm

import "testing"

func TestDelete(t *testing.T) {
	GetInstance().Where("name=?", "fangfang").Delete(&User{})
}
