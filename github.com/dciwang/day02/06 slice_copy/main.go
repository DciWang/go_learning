package main

import "fmt"

/*slice  copy()*/
func main() {
	s1 := []int{1,2,3}
	var s2=make([]int,2,3)
	s3 := s1     //赋值   s3跟s1指向同一个底层数组
	copy(s2,s1)   //将底层数组重新复制一份并指向
	s1[0]=4
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%v, cap(s2)=%v\n", s2, len(s2), cap(s2))
	fmt.Printf("s3=%v, len(s3)=%v, cap(s3)=%v\n", s3, len(s3), cap(s3))
	/*删除元素*/
	s4 :=append(s3[:1],s3[2:]...)
	fmt.Printf("s4=%v, len(s4)=%v, cap(s4)=%v\n", s4, len(s4), cap(s4))
	fmt.Println("s1",s1)    //append()  修改了底层的数组




}
