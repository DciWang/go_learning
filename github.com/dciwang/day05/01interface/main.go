package main

import "fmt"

/*接口是一种类型，它规定了变量有哪些方法*/

type cat struct {
	name string
}
type dog struct {
	name string
}
type person struct {
	name string
}

/*定义接口，里面写需要实现的方法*/
type speak interface {
	speak()
}

func (p person) speak() {
	fmt.Println("啊啊啊")
}
func (d dog) speak() {
	fmt.Println("汪汪汪")
}
func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func da(x speak) {
	//传进来一个参数，接收到的是什么我就打什么
	x.speak()
}

func main() {
	wdc := person{"wangdc"}
	wangcai := dog{"wangcai"}
	kafei := cat{"jiafei"}
	da(wdc)
	da(wangcai)
	da(kafei)
	var ss speak
	ss = wdc //类似多态
	ss.speak()
	ss = wangcai //类似多态
	ss.speak()
	ss = kafei //类似多态
	ss.speak()
}
