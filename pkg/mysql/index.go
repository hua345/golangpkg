package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

// root:pw@unix(/tmp/mysql.sock)/myDatabase?charset=utf8
// user:password@tcp(localhost:5555)/dbname?charset=utf8
// #TCP using default port (3306) on localhost:
// user:password@tcp/dbname?charset=utf8mb4,utf8&sys_var=esc%40ped

type MysqlConfig struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

var mysqlConfig = &MysqlConfig{
	UserName: "root",
	Password: "",
	Host:     "127.0.0.1",
	Port:     3306,
	Database: "db_example",
}

const db_URL = "springuser:123456@tcp(127.0.0.1:3306)/db_example?charset=utf8"

var (
	once sync.Once

	instance *sql.DB
)

func GetInstance() *sql.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host,
			mysqlConfig.Port, mysqlConfig.Database)
		var err error
		instance, err = sql.Open("mysql", dsn)
		if err != nil {
			fmt.Printf("Open mysql failed,err:%v\n", err)
			return
		}
		instance.SetMaxOpenConns(1024)
		instance.SetMaxIdleConns(16)
	})

	return instance
}

// In order to handle time.Time correctly, you need to include parseTime as a parameter.
// In order to fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4
func NewMysqlDb() {

}
