package main

import (
	"fmt"
	"os"
)

/*
函数版学生管理系统
	支持crud
*/

type student struct {
	id    int64
	name  string
	class string
}

func newStudent(id int64, name string, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}

}

//var students = []student{}
var (
	students = make(map[int64]*student)
)

func main() {
	for {
		//1：打印菜单
		fmt.Println("欢迎光临学生管理系统")
		fmt.Println(`
				1.查看所有学生
				2.新增学生
				3.删除学生
				4.退出
`)
		//2：等待用户选择要做什么
		fmt.Println("请输入你要干什么")
		var choose int
		fmt.Scanln(&choose)
		fmt.Printf("你选择了%d这个操作！\n", choose)
		//3:执行对应的函数
		switch choose {
		case 1:
			showAllStudent()
		case 2:
			insertStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("go off")
		}
	}

}

func showAllStudent() {
	for _, s := range students {
		fmt.Println("学生列表如下")
		fmt.Printf("学号：%d 姓名：%s ,班级：%s \n", s.id, s.name, s.class)
	}
}
func insertStudent() {
	var id int64
	var name string
	var classs string

	fmt.Println("请输入学生id")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入学生班级")
	fmt.Scanln(&classs)
	stu := newStudent(id, name, classs)

	students[stu.id] = stu
	fmt.Println("添加成功")
	showAllStudent()
}
func deleteStudent() {
	var id int64
	fmt.Println("请输入要删除的id")
	fmt.Scanln(&id)
	/*newS :=	&student{
		id: 0,
		name: "",
		class: "",
	}*/
	if students[id] == nil {
		fmt.Println("该学生不存在,重新输入")
		deleteStudent()
	} else {
		delete(students, id)
		fmt.Println("删除成功")
	}
}
