package main

import (
	"fmt"
	"os"
	"sort"
)

// 函数版学生管理系统
// 查看、新增、删除学生

type student struct {
	id   int
	name string
}

var stu = make(map[int]*student)

func title(txt string) {
	fmt.Printf("======== %s =======\n", txt)
}

func showAllStudent() {
	title("查看所有学生")
	if len(stu) <= 0 {
		fmt.Println("(当前没有学生数据)")
		return
	}
	keys := make([]int, 0, len(stu))
	for k := range stu {
		keys = append(keys, k)
	}
	sort.Ints(keys)
    fmt.Printf("学号\t姓名\n")
    fmt.Printf("----\t-----\n")
	for _, k := range keys {
		fmt.Printf("%d\t%s\n", stu[k].id, stu[k].name)
	}
}

func addStudent() {
	title("新增学生")
	var s student
	for {
		var id int
		fmt.Println("学号 姓名")
		fmt.Scanln(&id, &s.name)
		if _, ok := stu[id]; !ok {
			s.id = id
			stu[id] = &s
			fmt.Println("学生添加成功：", s)
			break
		}
		fmt.Printf("学号%v已存在，请重新输入", id)
	}
}

func deleteStudent() {
	title("删除学生")
	var id int
	fmt.Println("请输入要删除的学号")
	fmt.Scanln(&id)
	if s, ok := stu[id]; !ok {
		fmt.Println("不存在的学号")
	} else {
		delete(stu, id)
		fmt.Println("已删除学生：", s)
	}
}
func main() {
	for {
		// 1. 打印菜单
		title("欢迎使用学生管理系统")
		fmt.Println(`
    1. 查看所有学生
    2. 新增学生
    3. 删除学生
    4. 退出
    `)
		// 2. 等待用户选择操作
		var chioce int
		fmt.Print("请输入你要执行的操作：")
		fmt.Scanln(&chioce)
		// 3. 执行对应的函数
		switch chioce {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			fmt.Println("再见")
			os.Exit(0)
		}
	}
}
