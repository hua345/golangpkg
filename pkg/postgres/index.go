package postgres

import (
	"database/sql"
	"fmt"
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

var postgresDB *sql.DB

func NewPostgresDB() {
	// sslmode就是安全验证模式;
	// sslmode=disable
	// sslmode=verify-full
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", postgresConfig.UserName, postgresConfig.Password, postgresConfig.Host, postgresConfig.Port, postgresConfig.Database)
	var err error
	postgresDB, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	postgresDB.SetMaxOpenConns(1024)
	postgresDB.SetMaxIdleConns(16)
}
