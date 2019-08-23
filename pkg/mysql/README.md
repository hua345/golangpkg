# 参考

- [https://www.mysql.com/](https://www.mysql.com/)
- [https://github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

## 安装

```bash
go get -u github.com/go-sql-driver/mysql
```

## 使用

> Go MySQL Driver is an implementation of Go's database/sql/driver interface. 
> You only need to import the driver and can use the full database/sql API then.

Use mysql as driverName and a valid DSN as dataSourceName:

```
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")
```
