# go module

是go 1.11 版本之后，官方推出的版本管理工具，并且从 go 1.13 开始，`go module`成为go语言默认的依赖管理工具。

## `GO111MODULE`

要雇用 `go module` 支持，首先要设置环境言火日王`GO111MODULE`，通过它可以开启/关闭模块支持。它有三个可选值：

+ `GO111MODULE=off`：禁用，编译时会从 `GOPATH` 和 `vendor` 文件夹中查找包

+ `GO111MODULE=on`：启用，编译时会忽略 `GOPATH` 和 `vendor` 文件夹，只根据 `go.mod` 下载依赖。

+ `GO111MODULE=auto`(默认)：当项目在 `$GOPATH/src` 外，且项目根目录有 `go.mod` 文件时，开启模块支持。

简单来说，设置 `GO111MODULE=on` 之后，就可以使用 `go module` 了，以后就没必要在`GOPATH` 中创建项目了，并且还能够很好的管理项目依赖的第三方包信息。

使用 `go module` 管理依赖后，会在项目根目录下生成 `go.mod` 和 `go.sum` 两个文件。

## `GOPROXY`

设置方法：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

### `go mod` 命令

格式：`go mod <cmd>`，常用`<cmd>` 如下：

|`cmd`|说明|
|----|----|
|`download`|下载依赖的module到本地cache(默认为`$GOPATH/pkg/mod`目录)|
|`edit`|编辑 `go.mod` 文件|
|`graph`|打印模块依赖图|
|`init`|初始化当前文件夹，创建 `go.mod` 文件|
|`tidy`|增加缺少的module，删除无用的module|
|`vendor`|将依赖复制到`vendor`下|
|`verify`|校验依赖|
|`why`|解释为什么需要依赖|

### `go.mod` 文件

记录了项目所有的依赖信息

## 使用 `go module`

### 既有项目

+ 在项目目录下执行 `go mod init`，生成一个 `go.mod` 文件

+ 执行 `go get`，查找并记录当前项目的依赖，同时生成一个 `go.sum` 文件，记录每个依赖库的版本和哈希值

### 新项目

+  `go mod init 项目名`命令，生成一个 `go.mod` 文件

+ 手动编辑`go.mod` 中的 `require` 依赖项或执行 `go get` 自动发现、维护依赖。
