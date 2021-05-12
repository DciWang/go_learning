package main

import "fmt"

type Person struct {
	name   string
	gender string
}

func main() {
	/*创建指针类型结构体*/
	wdc := new(Person)
	(*wdc).name = "wdc"
	wdc.gender = "男"
	fmt.Printf("wdc'tpye:%T\n", wdc)     //wdc'tpye:*main.Person
	fmt.Printf("wdc'value-> %#v\n", wdc) //wdc'value-> &main.Person{name:"wdc", gender:"男"}
	/*取结构体地址实例化*/
	q := &Person{}
	(*q).name = "lq"
	q.gender = "girl"
	fmt.Printf("q'type -> %T\n", q) //q'type -> *main.Person       使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	/*结构体内存布局*/
	f()
}

type x struct {
	a int8
	b int8
	c int8
	d int8
}

func f() {
	n := &x{1, 2, 3, 4}

	fmt.Printf("n'point -> %v\n", n)
	//地址是连续的
	fmt.Printf("n.a'point -> %v\n", &n.a) //n.a'point -> 0xc000096010
	fmt.Printf("n.b'point -> %v\n", &n.b) //n.a'point -> 0xc000096011
	fmt.Printf("n.c'point -> %v\n", &n.c) //n.a'point -> 0xc000096012
	fmt.Printf("n.d'point -> %v\n", &n.d) //n.a'point -> 0xc000096013
}
