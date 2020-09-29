# 字符串

+ 原生数据类型

+ 使用 utf-8 实现

+ 使用双引号 `""`

```go
// 字符串
s := "Hello 沙河"

// 字符
c1 := 'h'
c2 := '1'
c3 := '沙'

// 字节(byte)，1字节 = 8 bit（8个二进制位）
// 1个字符'A' = 1字节
// 1个UTF-8编码的汉字'沙' = 一般占3个字节
```

## 多行字符串

使用 `\``

## 字符串的常用操作

|方法|说明|
|----|----|
|`len(str)`|求长度|
|`+`或`fmt.Sprintf`|拼接字符串|
|`strings.Split`|分割|
|`strings.contains`|判断是否包含|
|`strings.HasPrefix/HasSuffix`|判断前缀/后缀|
|`strings.Index/LastIndex`|子串出现的位置|
|`strings.Join([]string, string)`|将数组合并为字符串|

## byte和rune类型

组成每个字符串的元素叫“字符”，可以通过遍历或者单个获取字符串元素来获得字符。
Go语言的字符有两种：


