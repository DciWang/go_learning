package main

import "fmt"

type Person struct {
	name   string
	age    int64
	gender string
}

type Student struct {
	class  string
	name   string
	Person Person
}
type Son struct {
	class  string
	name   string
	Person //匿名嵌套结构体，.属性   的方式获取值，先从外面结构体找，再从内部结构体找
}

func main() {
	wdc := Student{
		class: "wuban",
		name:  "q",
		Person: Person{
			name:   "王多聪",
			age:    23,
			gender: "男",
		},
	}

	q := Son{
		class: "5",
		name:  "wdc",
		Person: Person{
			name:   "q",
			age:    34,
			gender: "女",
		},
	}
	fmt.Printf("wdc'value -> %v\n", wdc)
	fmt.Printf("q.name'value -> %v\n", q.name) //q.name'value -> wdc
}
