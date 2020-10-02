# 切片(slice)

因为数组的长度是固定的，并且数组长度属于类型的一部分，所以数组有很多局限性。例如：

```go
func arraySum(x [3]int) int {
    sum := 0
     for _, v := range x {
        sum +=v
     }
    return sum
}
```

这个求和函数只能接受 `[3]int` 类型，其它的都不支持。再比如：

```go
a := [3]int{1, 2, 3}
```

这个数组已经有3个元素了，无法再往数组中添加新元素。

## 切片

是一个拥有相同类型元素的*可变长度的序列*。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。它是一个**引用类型**，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。

### 切片的定义

声明切片类型的基本语法：

```go
var name []T
```

其中：

+ `name`：变量名

+ `T`：切片中元素的类型

例如：

```go
var a [] string //声明一个string切片
```

### 长度和容量

+ `len()`：求长度

+ `cap()`：求容量

切片的长度就是它的元素个数

切片的容易是底层数组从切片的第一个元素到数组最后一个元素的数量

### 基于数组定义切片

```go
a := [5]int {1, 2, 3, 4, 5}
s := a[1:4]
```
