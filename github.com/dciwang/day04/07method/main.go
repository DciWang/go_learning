package main

import (
	"fmt"
)

/*
标识符：变量名 类型名 方法名
Go语言中如果标识符首字母是大写的，就表示是对外可见的（公共的）
*/
// Dog 这是一个结构体
type Dog struct {
	name string
}

type cat struct {
	name string
}

type Person struct {
	name string
	age  int32
}

//方法是特定结构体调用的函数
//接受者是表示的是调用该方法的具体类型变量，多用的类型首字母小写
func (d Dog) f(name string) string {
	fmt.Println("Dog say  汪汪汪", d.name) //Dog say  汪汪汪 旺财
	return name + d.name
}

func (p Person) guonian() { //这里是值接受者
	p.age++
}

func (p *Person) guoNian() { //这里是指针接受者    一般使用指针接受者
	p.age++
}

/*给一个int 类型内置一个方法*/
type myInt int

func (m myInt) helloWorld() {
	fmt.Println("hello world")

}

func main() {
	d := Dog{"旺财"}
	c := cat{"加菲"}
	e := d.f(c.name)
	fmt.Println(e)
	fmt.Printf("c'value -> %#v\n", c) //c'value -> main.cat{name:"加菲"}
	//c.f()     编译不通过

	wdc := Person{"wdc", 18}
	fmt.Println("wdc'age-->", wdc.age)
	wdc.guonian()
	fmt.Println("wdc'age-->", wdc.age)
	wdc.guoNian()
	fmt.Println("wdc'age-->", wdc.age)

	m := myInt(10)
	m.helloWorld()

}
