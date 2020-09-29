package main

import "fmt"

func main() {
	for row := 1; row <= 9; row++ {
		for col := 1; col <= row; col++ {
			fmt.Printf("%d*%d=%d\t", col, row, row*col)
		}
		fmt.Println()
	}
}
