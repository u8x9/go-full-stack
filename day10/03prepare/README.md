# 预处理

+ 优化 mysql 服务器重复执行 sql 的方法，可以提升性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本

+ 避免 sql 注入

## go 的实现

```go
func (db *DB) Prepare(query string)(*Stmt, error)
```


