package main

import (
	"fmt"
)

type myInt int

type person struct {
	name string
	age  int
}

func main() {
	//问题1
	m := myInt(10) //强制类型转换
	fmt.Println("m -->", m)

	//问题2：结构体初始化

	//方法一：
	var p1 person
	p1.name = "wdc"
	p1.age = 18
	//方法2：
	//键值对初始化
	p2 := person{
		name: "wdc",
		age:  18,
	}
	fmt.Println("p2-->", p2)
	//值列表初始化
	p3 := person{"wdc", 18}
	fmt.Println("p3-->", p3)

}

/*问题3：为什么要有构造函数*/
func newPeson(name string, age int) *person {
	//别人调用我，我能给他返回一个person类型的变量
	return &person{
		name: name,
		age:  age,
	}

}
