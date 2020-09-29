讲师博客：https://www.liwenzhou.com/posts/Go/go_menu/

## 编译

使用`go build`

1. 在项目目录下执行`go build`
2. 在其它路径下执行`go build`，需要在后面加上项目的路径（GOPATH/src之后的路径）
3. `-o` 指定输出的可执行文件的名字：`go build -o foo.bar`

## 交叉编译

```bash
CGO_ENABLED=0 # 禁用CGO
GOOS=linux # 目标平台是linux
GOARCH=amd64 # 架构
```

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=darin GOARCH=amd64 go build
```
