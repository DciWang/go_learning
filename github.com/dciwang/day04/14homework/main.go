package main

import (
	"fmt"
	"os"
)

/**
对象版学生管理系统
	1.有一个管理者
		保存了一些数据    --->结构体字段
		有四个功能    ---->结构体方法
*/
type Student struct {
	id    int32
	name  string
	class string
}

func newStudent(id int32, name string, class string) *Student {
	return &Student{
		id:    id,
		name:  name,
		class: class,
	}
}

type StudentMgr struct {
	allStudent map[int32]*Student
}

func (s StudentMgr) showAllStudent() {
	for _, s := range s.allStudent {
		fmt.Println("学生列表如下")
		fmt.Printf("学号：%d 姓名：%s ,班级：%s \n", s.id, s.name, s.class)
	}
}
func (s StudentMgr) insertStudent() {
	var id int32
	var name string
	var classs string

	fmt.Println("请输入学生id")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入学生班级")
	fmt.Scanln(&classs)
	stu := newStudent(id, name, classs)

	s.allStudent[stu.id] = stu
	fmt.Println("添加成功")
	s.showAllStudent()
}
func (s StudentMgr) deleteStudent() {
	var id int32
	fmt.Println("请输入要删除的id")
	fmt.Scanln(&id)
	/*newS :=	&student{
		id: 0,
		name: "",
		class: "",
	}*/
	if s.allStudent[id] == nil {
		fmt.Println("该学生不存在,重新输入")
		s.deleteStudent()
	} else {
		delete(s.allStudent, id)
		fmt.Println("删除成功")
	}
}
func (s StudentMgr) updateStudent() {
	var id int32
	fmt.Println("请输入要修改学生的id")
	fmt.Scanln(&id)
	/*newS :=	&student{
		id: 0,
		name: "",
		class: "",
	}*/
	if s.allStudent[id] == nil {
		fmt.Println("该学生不存在,重新输入")
		s.deleteStudent()
	} else {
		s.insertStudent()
		fmt.Println("修改成功")
	}
}

func main() {

	for {
		var sm StudentMgr
		fmt.Println("欢迎光临学生管理系统")

		fmt.Println(`
				1.查看所有学生
				2.新增学生
				3.删除学生
				4.退出
`)
		fmt.Println("请输入你要干什么")
		var choos int32
		fmt.Scanln(&choos)

		switch choos {
		case 1:
			sm.showAllStudent()
		case 2:
			sm.insertStudent()
		case 3:
			sm.updateStudent()
		case 4:
			sm.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("go off")
		}
	}

}
