# 并发

[讲师博客](https://www.liwenzhou.com/posts/Go/14_concurrence/?from=u8x9@github)

## 并发与并行

+ 并发：同一时间段内执行多个任务

+ 并行：同一时刻执行多个任务

go 语言通过 `goroutine` 实现。它类似线程，属于用户态的线程。`goroutine` 由go语言的runtime调度完成，而线程由操作系统调度完成。

go还提供`channel`在多个`goroutine`间进行通信。`goroutine` 和 `channel` 是go语言秉承 csp 并发模式的重要实现基础。
