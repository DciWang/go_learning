package main

import "fmt"

/**
切片就是一个框，框住了一块连续的内存，所以存储的数据类型都必须是一样的
切片不能相互比较，只能跟nil比较

*/
func main() {
	//make函数创建切片
	s1 := make([]int, 5, 10) //len(s1)=5,cap(s1)=10
	fmt.Printf("s1:%v len,(s1):%d,cap(s1):%d\n", s1, len(s1), cap(s1))
	/*切片的拷贝*/

	s2 := []int{1, 2, 3}
	s3 := s2 //切片是引用类型，都指向底层的数组
	fmt.Printf("s2:%v,s3:%v\n", s2, s3)
	s2[2] = 100
	fmt.Printf("s2:%v,s3:%v\n", s2, s3)
	fmt.Println(len(s2) == 0) //判断切片是否为空
	/*遍历切片*/
	//1.索引遍历
	for i := 0; i < len(s2); i++ {
		fmt.Printf("i:%v\n", s2[i])
	}
	//2.for rang循环
	for i, i2 := range s2 {
		fmt.Println(i, i2)
	}

}
