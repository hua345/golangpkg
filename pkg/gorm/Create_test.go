package gorm

import (
	"golangpkg/pkg/redis"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	NewGorm()
	redis.NewRedis()
	leaf := redis.NewLeaf("bookKey")
	bookId := leaf.NextId()

	user := &User{
		Name:       "fangfang",
		UserMobile: "17875266970",
		Age:        24,
	}
	user.ID = bookId
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	t.Log(user.CreatedTime.Format("2006/01/02 15:04:05"))
	gormDB.Create(&user)
}
