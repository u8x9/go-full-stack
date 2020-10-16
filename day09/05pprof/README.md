# Go 性能优化

Go 语言项目中的性能优化主要有以下几个方面：

+ **CPU Profile**：报告程序的 cpu 使用情况，按照一定频率去采集应用程序在 cpu 和寄存器上面的数据

+ **Memory Profile(Heap Profile)**：报告程序的内存使用情况

+ **Block Profile**：报告 goroutines 不在运行状态的情况，可以用来分析和查找死锁等性能瓶颈

+ **Goroutine Profile**：报告 goroutines 的使用情况，有哪些 goroutine，它们的调用关系是怎么样的

## 采集性能数据

标准库：

+ `runtime/pprof`：采集工具型应用运行数据进行分析

+ `net/http/pprof`：采集服务型应用运行时数据进行分析

pprof开启后，每隔一段时间(10ms)就会收集一次当前的堆栈信息，获取个个函数占用的cpu以及内存资源。最后通过这些采样数据进行分析，形成一个性能分析报告。

> 注意：只应在性能测试的时候才在代码中引入pprof

## 工具型应用

运行一段时间就会结束退出的应用程序。这类应用最好的办法是在应用退出的时候把 profiling 的报告保存到文件中进行分析。可以使用 `runtime/pprof` 库。

### cpu 性能分析

开启 cpu 性能分析

```go
pprof.StartCPUProfile(w io.Writer)
```

停止 cpu 性能分析

```go
pprof.StopCPUProfile()
```

程序结束后，就会生成一个文件，存储了 CPU Profiling 数据。得到采样数据之后，使用 `go tool pprof` 工具进行 cpu 性能分析。

### 内存性能优化

记录程序的堆栈信息

```go
pprof.WriteHeapProfile(w io.Writer)
```

得到采样数据之后，使用 `go tool pprof` 工具进行内存性能分析。

`go tool pprof` 默认使用 `-inuse_space` 进行统计，还可以使用 `-inuse-objects` 查看分析对象的数量。

## 服务型应用

如果使用了默认的 `http.DefaultServeMux`（就是直接使用 `http.ListenAndServe(":8080", nil)`），只需要在 web server 端代码中按如下方式导入：

```go
import _ "net/http/pprof"
```

如果使用了自定义的Mux，则需要手动注册一些路由规则

```go
r.HandleFunc("/debug/pprof/", pprof.Index)
r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
r.HandleFunc("/debug/pprof/profile", pprof.Profile)
r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
r.HandleFunc("/debug/pprof/trace", pprof.Trace)
```

如果使用的是 gin，推荐使用github.com/gin-contrib/pprof，在代码中通过以下命令注册pprof相关路由。

```go
pprof.Register(router)
```

无论哪种方式，HTTP 服务都会多出`/debug/pprof`。该路径下还有几个子页面：

+ `/debug/pprof/profile`：访问这个链接会自动进行 CPU profiling，持续 30s，并生成一个文件供下载

+ `/debug/pprof/heap`： Memory Profiling 的路径，访问这个链接会得到一个内存 Profiling 结果的文件

+ `/debug/pprof/block`：block Profiling 的路径

+ `/debug/pprof/goroutines`：运行的 goroutines 列表，以及调用关系

## `go tool pprof`

```bash
go tool pprof [binary] [source]
```

+ `binary` 应用的二进制文件，用来解析各种符号

+ `source` 表示 profile 数据的来源，可以是本地文件也可以是http地址

