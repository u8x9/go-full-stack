package main

import "fmt"

// 回文判断
// 上海自来水来自海上
// 山西运煤车煤运西山
// 黄山落叶松叶落山黄

func main() {
	s := "上海自来水来自海上"
	ss := []rune(s)
	ssLen := len(ss)
	ok := true
	for i := 0; i < ssLen/2; i++ {
		if ss[i] != ss[ssLen-1-i] {
			ok = false
			break
		}
	}
	fmt.Printf("%q 是回文？%v\n", s, ok)
}
