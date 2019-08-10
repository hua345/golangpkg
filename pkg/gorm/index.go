package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

var postgresConfig = &PostgresConfig{
	UserName: "fang",
	Password: "123456",
	Host:     "192.168.137.128",
	Port:     5432,
	Database: "fangdb",
}

var gormDB *gorm.DB

func NewGorm() {
	// sslmode就是安全验证模式;
	// sslmode=disable
	// sslmode=verify-full
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", postgresConfig.UserName, postgresConfig.Password, postgresConfig.Host, postgresConfig.Port, postgresConfig.Database)
	var err error
	gormDB, err = gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("Open gorm failed,err:%v\n", err)
		return
	}
	gormDB.DB().SetMaxIdleConns(16)
	gormDB.DB().SetMaxOpenConns(128)
	// 启用Logger，显示详细日志
	gormDB.LogMode(true)
	// 全局禁用表名复数
	gormDB.SingularTable(true)
}
