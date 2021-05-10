package main

import "fmt"

func main() {
	var p person
	p.name="wdc"
	p.gender="男"
	ff(p)
	fmt.Println(p.gender)
	f(&p)
}
/*结构体是值类型*/
type person struct {
	name,gender string
}
func f(p *person)  {
	(*p).gender="不男不女"

}

func ff(p person)  {
	p.gender="女"    //修改的是p的副本，值传递
	fmt.Println(p.gender)   //女
}