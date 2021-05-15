package main

import "fmt"

//动态数组，长度可变

/**
切片指向了一个底层的数组
切片的长度是它元素的个数
切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
*/
func main() {
	//定义
	var i1 []int
	var i2 []int
	//初始化
	var s1 = []string{}
	s2 := []string{}
	i1 = []int{1, 2, 3}
	s1 = []string{"甘肃", "武威", "民勤"}
	fmt.Println(i2, s2)
	fmt.Println(i2 == nil) //true  只定义但是没有开辟空间
	fmt.Println(s2 == nil) //false   开辟了空间
	fmt.Println(s1, i1)
	fmt.Println(s1 == nil)
	fmt.Println(i1 == nil)

	/*切片的长度len()与容量cap()*/
	fmt.Printf("len(i1):%v,cap(i1):%v\n", len(i1), cap(i1))
	fmt.Printf("len(s1):%v,cap(s1):%v\n", len(s1), cap(s1))

	//1、由数组得到切片
	v1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	sl := v1[1:5] //基于一个数组得到切片[3, 5, 7, 9],左闭右开
	fmt.Println("sl=", sl)
	sl1 := v1[:4] //从0开始到3结束  ====> [0:3]
	sl2 := v1[3:] //从3开始到len()结束    ====>[3:len(s1)]  ------> [7, 9, 11, 13]
	fmt.Println("sl1=", sl1)
	fmt.Println("sl2=", sl2)
	//切片的容量是底层数组从切片的第一个元素到数组的最后一个数组的数量
	fmt.Printf("len(sl1):%v,cap(sl1):%v\n", len(sl1), cap(sl1))
	fmt.Printf("len(sl2):%v,cap(sl2):%v\n", len(sl2), cap(sl2))

	//2、切片再切片
	sl3 := sl2[3:]
	fmt.Println("sl3=", sl3)
	fmt.Printf("len(sl3):%v,cap(sl3):%v\n", len(sl3), cap(sl3)) //在数组 [7, 9, 11, 13] 上切片，得到的结果为[13]

	/*切片是引用类型，会根据指向数组的值改变而改变，数组是值类型的值*/
	fmt.Println("sl2=", sl2) //[7 9 11 13]
	v1[6] = 1300
	fmt.Println("sl2=", sl2) //[7 9 11 1300]

}
