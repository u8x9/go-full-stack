package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(dst, src string) (int64, error) {
	fsrc, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer fsrc.Close()
	fdst, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return 0, err
	}
	defer fdst.Close()
	const bufSize = 1024
	var buf = make([]byte, bufSize)
	var length int64 = 0

	for {
		n, err := fsrc.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return length, err
		}
		length += int64(n)
		_, err = fdst.Write(buf[:n])
		if err != nil {
			return length, err
		}
	}
	return length, nil
}

func copyFileByAPI(dst, src string) (int64, error) {
	fsrc, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer fsrc.Close()
	fdst, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return 0, err
	}
	defer fdst.Close()
	return io.Copy(fdst, fsrc)
}
func main() {
	//n,err:=copyFile("main.go.copy", "main.go")
	n, err := copyFileByAPI("main.go.copy", "main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("成功拷贝：", n)
}
