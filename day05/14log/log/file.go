package log

import "os"

func GetFileHandler(filename string)(*os.File, error) {
    return os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}

func MustGetFileHandler(filename string) *os.File {
    f, err := GetFileHandler(filename)
    if err!=nil{
        panic(err)
    }
    return f
}
