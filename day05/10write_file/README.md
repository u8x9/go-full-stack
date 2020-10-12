# 文件写入操作

## `os.OpenFile()`

```go
os.OpenFile(name string, flag int, perm FileMode) (*os.File, error)
```

+ `name`：要打开的文件名

+ `flag`：打开文件的模式

|模式|说明|
|----|----|
|`os.O_WRONLY`|只写|
|`os.O_CREATE`|当文件不存在时，创建文件|
|`os.O_RDONLY`|只读|
|`os.O_RDWR`|读写|
|`os.O_TRUNC`|清空文件原有内容|
|`os.O_APPEND`|追加|

+ `perm`：权限，八进制数
    - `04`：(r)读
    - `02`：(w)写
    - `01`：(x)可执行

