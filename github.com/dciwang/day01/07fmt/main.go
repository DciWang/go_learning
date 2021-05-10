package main

import "fmt"

func main() {
	n :=100
	fmt.Printf("%T\n",n)  //查看类型
	fmt.Printf("%v\n",n)  //查看值
	fmt.Printf("%b\n",n)  //查看二进制值
	fmt.Printf("%d\n",n)  //查看十进制值
	fmt.Printf("%o\n",n)  //查看八进制值
	fmt.Printf("%x\n",n)  //查看十六进制值
	fmt.Printf("%d%%\n",n)  //百分树，用%来转义%
	x :=false

	fmt.Printf("%t\n",x)     //查看布尔类型
	m :="hello 北京"
	fmt.Printf("%s\n",m)  //查看字符串值
	fmt.Printf("%v\n",m)  //查看值
	fmt.Printf("%p\n",&m)  //查看指针
	fmt.Printf("%#v\n",m)  //查看值，并显示详细信息，如查看字符串值，会加引号
	for _, i := range m {
		fmt.Printf("%c\n",i)  //查看字符
	}

}
