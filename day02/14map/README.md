# map

是一种无序的，基于`key-vale`的数据结构，它是引用类型，必须初始化才能使用

## map 定义

语法：

```go
map[KeyType]ValueType
```

+ `KeyType`：键的类型

+ `ValueType`：值的类型

map 变量默认初始值为 `nil`，需要使用 `make()` 来分配内存：

```go
make(map[KeyType]ValueType, cap)
// 或者
make(map[KeyType]ValueType)
```
