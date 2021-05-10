package main

import "fmt"

func main() {
/*	var  s string
	fmt.Scan(&s)
	fmt.Println("用户输入的内容是：",s)*/

	var (
		name string
		age int
		class string
	)

	fmt.Scanf("%s %d %s\n",&name,&age,&class)
	fmt.Printf("%s %d %s\n",name,age,class)

}
