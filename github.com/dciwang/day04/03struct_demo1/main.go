package main

import "fmt"

func main() {
	fmt.Println(f().name)
	fmt.Println(f().age)
	fmt.Println(f().hobby)
	fmt.Printf("%T\n",f())
}

type user struct {
	id  uint64
	age uint64
	name string
	hobby []string
}

func f() user{
	var wdc user
	wdc .age =1
	wdc .name="wind"
	wdc.hobby=[]string{"basketball","football"}

	/**
	匿名结构体
	 */
	var s struct{
		x string
		y int
	}
	s.x="aaa"
	s.y=20
	fmt.Printf("type:%T value:%v\n",s,s)
	return wdc

}
