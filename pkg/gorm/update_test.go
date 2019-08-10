package gorm

import "testing"

func TestUpdate(t *testing.T) {
	NewGorm()

	var user User
	// Update single attribute if it is changed
	gormDB.Model(&user).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;
	t.Log(user)
}
