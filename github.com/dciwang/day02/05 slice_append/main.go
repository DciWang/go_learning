package main

import "fmt"

/*append()	为切片添加元素*/
func main() {
	s1 := []string{"上海", "云南", "杭州", "成都", "重庆"}
	//s1[5] = "民勤"    发生越界异常

	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1))   //容量为5

	s1 =append(s1,"民勤") //调用append()函数必须用原来的切片变量接受返回值，append()追加元素，原来的底层数组放不下时，Go底层就会把底层数组换一个
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1))    //容量扩容为10
	s2 := []string{"深圳"}
	s1 = append(s1, s2...)  //...表示拆开
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1))



}
