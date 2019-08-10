package gorm

import "testing"

func TestUpdate(t *testing.T) {
	NewGorm()

	// Update single attribute if it is changed
	gormDB.Model(User{}).Update("name", "hello")
	// UPDATE users SET name='hello';

	// Update single attribute with combined conditions
	gormDB.Model(User{}).Where("age >= ?", 20).Update("name", "fangfang")

}
