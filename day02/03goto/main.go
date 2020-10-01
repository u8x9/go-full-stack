package main

import "fmt"

func main() {
	//flag := false
	//for i := 0; i < 10; i++ {
	//for j := 'A'; j < 'Z'; j++ {
	//flag = j == 'C'
	//if flag {
	//break
	//}
	//fmt.Printf("%d %c\n", i, j)
	//}
	//if flag {
	//break
	//}
	//}
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto breakFor
			}
			fmt.Printf("%d %c\n", i, j)
		}
	}
breakFor:
	fmt.Println("over")
}
