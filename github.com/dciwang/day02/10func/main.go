package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		num1 =rand.Int()
		num2 =rand.Int()
	)

	num3 :=f(num1,num2)
	fmt.Println(num3)
	fmt.Println(num2)
	fmt.Println(num1)
	f5()

}
/*返回值可以命名也可以不命名*/
//当返回值已经再返回值列表中定义的时候return可以省略
func f(int1 int,int2 int)(ret int) {
	ret = int1 + int2
	return
}
//当返回值再返回值列表中没有定义的话不能省略
func f2(int1 int) int  {
	ret := int1
	return  ret
}

//多个返回值,当多个参数的类型一样时，前面的参数类型可以省略
func f3(int1 ,int2 int,s1,s2 string,t,b bool)(int,int){
	sum := int1+int2
	ret := int1 * int2
	return  ret,sum
}

//可变长参数
//可变长参数，必须放在参数列表的最后
func f4(s string,i... int){
	fmt.Println(s)
	fmt.Println(i)   //类型是切片
}

func f5(){
	f4("aaa",2,3,4,45,5,5,)

}





