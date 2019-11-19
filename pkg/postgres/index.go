package postgres

import (
	"database/sql"
	"fmt"
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

	instance *sql.DB
)

func GetInstance() *sql.DB {
	once.Do(func() {
		// sslmode就是安全验证模式;
		// sslmode=disable
		// sslmode=verify-full
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			postgresConfig.UserName, postgresConfig.Password, postgresConfig.Host, postgresConfig.Port, postgresConfig.Database)
		var err error
		instance, err = sql.Open("postgres", dsn)
		if err != nil {
			fmt.Printf("Open mysql failed,err:%v\n", err)
			return
		}
		instance.SetMaxOpenConns(1024)
		instance.SetMaxIdleConns(16)
	})

	return instance
}
