package gorm

import (
	"github.com/hua345/golangpkg/pkg/redis"
	"testing"
	"time"
)

func TestTransaction(t *testing.T) {
	leaf := redis.NewLeaf("bookKey")
	bookId := leaf.NextId()
	user := &User{
		Name:       "transaction",
		UserMobile: "17875266970",
		Age:        24,
	}
	user.ID = bookId
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	// 开启事务
	tx := GetInstance().Begin()
	defer func() {
		if r := recover(); r != nil {
			// 如果发生错误则执行回滚
			tx.Rollback()
		}
	}()
	if err := tx.Create(user).Error; err != nil {
		// 如果发生错误则执行回滚
		tx.Rollback()
		return
	}
	// 在事务中执行具体的数据库操作 (事务内的操作使用 'tx' 执行，而不是 'db')
	if err := tx.Create(&User{Name: "Giraffe"}).Error; err != nil {
		// 如果发生错误则执行回滚
		tx.Rollback()
		return
	}
	tx.Commit()
}
