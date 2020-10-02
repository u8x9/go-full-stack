# 指针

+ `&`：取地址

+ `*`：根据地址取值

## `make` 和 `new` 的区别

+ `make` 和 `new` 都是用来分配内存

+ `new` 用来给非引用类型申请内存，返回对应类型的指针

+ `make` 也是用于内存分配，区别于`new`，它只用于slice, map以及chan的内存创建, 而且它返回的类型就是这三个类型本身，而不是它们的指针, 因为这三个类型本身就是引用类型，所以就没必要再返回它们的指针了
