package main

import "fmt"

const (
    _ = iota
    KB = 1 << (10 * iota)
    MB
    GB
    TB
    PB
)

func main(){
    fmt.Println("KB:", KB)
    fmt.Println("MB:", MB)
    fmt.Println("GB:", GB)
    fmt.Println("TB:", TB)
    fmt.Println("PB:", PB)
}
