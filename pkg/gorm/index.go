package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"sync"
)

type PostgresConfig struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

var postgresConfig = &PostgresConfig{
	UserName: "postgres",
	Password: "",
	Host:     "127.0.0.1",
	Port:     5432,
	Database: "db_example",
}

var (
	once sync.Once

	instance *gorm.DB
)

func GetInstance() *gorm.DB {
	once.Do(func() {
		// sslmode就是安全验证模式;
		// sslmode=disable
		// sslmode=verify-full
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", postgresConfig.UserName, postgresConfig.Password, postgresConfig.Host, postgresConfig.Port, postgresConfig.Database)
		var err error
		instance, err = gorm.Open("postgres", dsn)
		if err != nil {
			fmt.Printf("Open gorm failed,err:%v\n", err)
			return
		}
		instance.DB().SetMaxIdleConns(16)
		instance.DB().SetMaxOpenConns(128)
		// 启用Logger，显示详细日志
		//gormDB.LogMode(true)
		// 全局禁用表名复数
		instance.SingularTable(true)
	})

	return instance
}
