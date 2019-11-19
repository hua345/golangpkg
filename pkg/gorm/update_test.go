package gorm

import "testing"

func TestUpdate(t *testing.T) {
	// Update single attribute if it is changed
	GetInstance().Model(User{}).Update("name", "hello")
	// UPDATE users SET name='hello';

	// Update single attribute with combined conditions
	GetInstance().Model(User{}).Where("age >= ?", 20).Update("name", "fangfang")

}
