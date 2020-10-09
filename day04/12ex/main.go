package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int
	name string
}

func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

type manager struct {
	stu map[int]*student
}

func newManager() *manager {
	return &manager{
		stu: make(map[int]*student),
	}
}

func (m *manager) showAllStudent() {
	for k, v := range m.stu {
		fmt.Println(k, v.name)
	}
}

func (m *manager) getStudent(id int) *student {
	if s, ok := m.stu[id]; ok {
		return s
	}
	return nil
}

func (m *manager) addStudent() {
	var (
		id   int
		name string
	)
	fmt.Println("学号 姓名")
	fmt.Scanln(&id, &name)
	s := m.getStudent(id)
	if s != nil {
		fmt.Println("已存在：", s)
		return
	}
	s = newStudent(id, name)
	m.stu[id] = s
	fmt.Println("添加成功：", s)
}
func (m *manager) deleteStudent() {
	var id int
	fmt.Println("请输入要删除的学号")
	fmt.Scanln(&id)
	s := m.getStudent(id)
	if s == nil {
		fmt.Println("查无此人")
		return
	}
	delete(m.stu, id)
	fmt.Println("删除成功：", s)
}

func main() {
	m := newManager()
	for {
		fmt.Println(`
    1. 所有学生
    2. 添加学生
    3. 删除学生
    4. 退出
    `)
    fmt.Print("请选择你要的操作：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			m.showAllStudent()
		case 2:
			m.addStudent()
		case 3:
			m.deleteStudent()
		case 4:
			fmt.Println("再见")
			os.Exit(0)
		}
	}
}
