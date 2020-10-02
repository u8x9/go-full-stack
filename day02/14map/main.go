package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["理想人生"] = 9000
	m1["古利哪扎"] = 7000
	m1["迪丽乐八"] = 8000
	m1["沙四笔呀"] = 6323
	fmt.Println(m1)
	fmt.Println(m1["NOT_EXISTS"]) // 如果不存在这个key, 拿到的是对应类型的零值

	// 健壮的写法
	if v, ok := m1["NOT_EXISTS"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此key")
	}

	// map 的遍历
	// 键值对
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除
	delete(m1, "理想人生")
	fmt.Println(m1)
	delete(m1, "NOT_EXISTS") // 删除不存在的key
}
