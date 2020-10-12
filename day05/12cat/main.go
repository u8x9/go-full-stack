package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func cat(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("useage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}
	if err := cat(os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
