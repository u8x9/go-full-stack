# MYSQL

## 安装驱动

```bash
go get -u github.com/go-sql-driver/mysql
```

## 使用

```go
import "database/sql"
_ import "github.com/go-sql-driver/mysql"
```

### `database/sql`

***原生支持连接池，是并发安全的。***

> `sql.DB` 是一个数据库(操作)句柄，代表一个具有零到多个底层连接的连接池， 它可以安全的被多个 goroutine 同时使用
