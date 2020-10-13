package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1(){
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("can not get runtime info")
		return
	}
	fmt.Printf("pc: %#v\n", pc)
	fmt.Println("file: ", file)
	fmt.Println("line: ", line)
    funName:=runtime.FuncForPC(pc).Name()
    fmt.Println("fun name: ", funName)
    fmt.Println("file name: ", path.Base(file))
}

func main() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("can not get runtime info")
		return
	}
	fmt.Printf("pc: %#v\n", pc)
	fmt.Println("file: ", file)
	fmt.Println("line: ", line)
    funName:=runtime.FuncForPC(pc).Name()
    fmt.Println("fun name: ", funName)
    fmt.Println("file name: ", path.Base(file))

    fmt.Println("-------")
    f1()
}
