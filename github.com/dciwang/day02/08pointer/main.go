package main

import "fmt"

/*
go 语言中不存在指针操作，但是要记住两个操作
1.&:取地址
2.*：根据地址取值
*/
func main() {

	 a := 2
	 fmt.Println(&a)   // 0xc0000a2000    16进制
	 b := &a
	 fmt.Printf("%T\n",b)    //*int   表示int类型的指针
	 fmt.Println(*b)  //2


	 /*定义一个内存地址*/
	 var a1 *int   //此时的 a1 为 nil
	fmt.Println(a1)
	//	*a1=10   //*a1  相当于对a1的对应的值赋值，a1为nil没有对应的值，报空指针异常
	/*所以一般直接用new()定义并初始化一个指针*/

	a2 := new(int)    //定义并初始化一个一个内存空间
	fmt.Println(a2)
	*a2 = 10
	fmt.Println(*a2)

	/*
注：new与make的区别，
	1。二者都是用来分配内存的
	2。make是用来初始化slice，make，chan。返回的还是这三个应用类型的本身
	3。new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
	*/


}
