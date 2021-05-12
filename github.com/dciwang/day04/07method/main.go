package main

import "fmt"

type dog struct {
	name string
}

type cat struct {
	name string
}

//方法是特定结构体调用的函数
//接受者是表示的是调用该方法的具体类型变量，多用的类型首字母小写
func (d dog) f() {
	fmt.Println("dog say  汪汪汪", d.name) //dog say  汪汪汪 旺财
}
func main() {
	d := dog{"旺财"}
	c := cat{"加菲"}
	d.f()
	fmt.Printf("c'value -> %#v\n", c) //c'value -> main.cat{name:"加菲"}
	//c.f()     编译不通过
}
