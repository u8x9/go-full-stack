# 同步、异步

+ `goroutine` 机制可以方便地实现异步处理

+ 另外，**在启动新的 `goroutine` 时，*不应该*使用原始上下文，*必须*使用它的*只读副本***