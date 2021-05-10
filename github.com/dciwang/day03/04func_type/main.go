package main

import "fmt"

type calculation func(int,int) int
func main() {
	var  c calculation  //创建一个calculation类型的函数
	c =aaa     //将aaa赋值给c
	fmt.Println("c=",c(1,2))
	fmt.Printf("%T\n",aaa)   //===>func(int, int) int
	fmt.Printf("%T\n",b)    //===>func(int, int) int
	fmt.Printf("%T\n",c)    //  ===>main.calculation

	ddd(aaa)
}

func aaa(x,y int)int {
	return  x+y
}
func b(x,y int)int {
	return  x-y
}

/*函数可以作为参数*/
func ddd(x func(s1,s2 int) int)  {   //函数ddd的参数为参数为string类型的s1，s2返回值为int类型的函数，ddd的返回值为int类型

}


/*函数可以作为返回值*/
func eee()func(x,y int) int {
	return aaa
}