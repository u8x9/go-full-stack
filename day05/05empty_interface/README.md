# 空接口

没有必要起名字，通常定义成下面的格式：

```go
interface{}
```

所有类型都实现了空接口，也就是说空接口可以用于存储任何类型

## 类型断言

```go
x.(T)
```

+ `x`：类型为 `interface{}` 的变量

+ `T`：`x`可能是的类型

```go
var xx interface{} = "Hello"
if v, ok := xx.(string); ok {
    fmt.Println(strings.ToUpper(v))
}
```
