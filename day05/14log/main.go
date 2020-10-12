package main

import (
	"os"

	"github.com/u8x9/go-full-stack/day05/14log/log"
)

func main() {
	l := log.NewLogger(os.Stdout)
	l.Debug("Hello")
	l.Info("世界")

	//f,err:=log.GetFileHandler("foo.log")
	//if err!=nil {
	//fmt.Println(err)
	//}
	f := log.MustGetFileHandler("foo.log")
	defer f.Close()
	l = log.NewLogger(f)
	l.Warn("你好")
	l.Error("World")
}
