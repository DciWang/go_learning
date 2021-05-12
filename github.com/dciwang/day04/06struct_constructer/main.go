package main

import "fmt"

type person struct {
	name   string
	gender string
}

/*
构造函数，约定成俗用new开头
返回的是结构体还是指针，根据现实情况决定
结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
*/
func newPerson(name, gender string) *person {
	return &person{
		name:   name,
		gender: gender,
	}
}

func main() {
	wdc := newPerson("wdc", "boy")
	fmt.Printf("wdc'value -> %#v\n", wdc)
}
