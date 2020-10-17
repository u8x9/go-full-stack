# Context

+ 推荐以参数的方式显式传递Context

+ 以Context作为参数的函数/方法，应该把Context作为第一个参数

+ 给一个函数/方法传递Context时，不要传递 `nil`。如果不知道传递什么，就使用`context.TODO()`

+ Context的`Value`相关方法应该传递请求域的必要数据，不应该用于传递可选参数

+ Context是线程安全的，可以放心的在多个 goroutine 中传递
