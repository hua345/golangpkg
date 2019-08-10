package gorm

import "testing"

func TestSelect(t *testing.T) {
	NewGorm()

	var user User
	gormDB.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	t.Log(user)

	// 通过主键查询最后一条记录
	user = User{}
	gormDB.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1
	t.Log(user)

	// Get first matched record
	user = User{}
	gormDB.Where("name = $1", "fangfang").First(&user)
	// SELECT * FROM users WHERE name = 'fangfang' limit 1;
	t.Log(user)
	// LIKE
	user = User{}
	gormDB.Where("name LIKE ?", "%fang%").First(&user)
	t.Log(user)
	// AND
	user = User{}
	gormDB.Where("name = ? AND age >= ?", "fangfang", "22").Find(&user)
	t.Log(user)
	// Struct
	user = User{}
	gormDB.Where(&User{Name: "fang", Age: 24}).First(&user)
	// SELECT * FROM users WHERE name = "fang" AND age = 24 LIMIT 1;
	t.Log(user)
}
