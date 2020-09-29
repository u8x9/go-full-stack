package main

import (
	"fmt"
	"strings"
)

func main(){
    // 多行字符串/原样输出
    s:=`日照香炉生紫烟
    遥看瀑布挂前川
    不识庐山真面目
    只缘身在此山中
    `
    fmt.Println(s)

    // 字符串常用操作
    // 长度
    slen := len(s)
    fmt.Println("s的长度是：", slen)

    // 拼接
    s1 := "Hello"
    s2 := s1 + "世界"
    s3 := fmt.Sprintf("%s of %d!!!", s2, 2020)
    fmt.Println(s3)

    // 分割
    s4 := "/etc/apt/apt.conf.d"
    arr := strings.Split(s4, "/")
    fmt.Println(arr)

    // 包含
    fmt.Println(strings.Contains(s3, "世界"))

    // 前缀/后缀
    fmt.Println(strings.HasPrefix(s4, "/"), strings.HasSuffix(s4, ".d"))

    // 位置
    idx := strings.Index(s4, "apt")
    lidx := strings.LastIndex(s4, "apt")
    fmt.Println("idx:", idx, ", lidx:", lidx)

    // join
    fmt.Println(strings.Join(arr, "+"))
}
