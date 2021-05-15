package main

import "fmt"

/*匿名字段，默认把类型当做名字，所以相同类型只能写一个一般不使用*/
type Person1 struct {
	string
	int64
}

func main() {
	p1 := Person1{"wangduocong", 26}
	fmt.Printf("p1'string -> %s\n", p1.string) //p1'string -> wangduocong

}
