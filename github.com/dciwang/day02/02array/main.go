package main

import "fmt"

func main() {

	//定义数组
	//var a [3]int
	//var b [4]int

	/*数组赋值*/
	var testArray [3]int                     //数组会初始化为int类型的零值
	numArray := [3]int{1, 2}                 //使用指定的初始值完成初始化
	cityArray := [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                   //[0 0 0]
	fmt.Println(numArray)                    //[1 2 0]
	fmt.Println(cityArray)                   //[北京 上海 深圳]

	c := testArray == numArray
	fmt.Println(c)
	/*数组遍历*/

	fmt.Println("for 遍历")
	cit := [...]string{"北京", "上海", "民勤"}
	for i := 0; i < len(cit); i++ {
		fmt.Println(cit[i])
	}

	fmt.Println("for range  遍历")
	for i, j := range cit {
		fmt.Println(i, j)
	}
	/*二维数组*/
	aa := [...][3]string{
		{"aaa", "aaaa", "fjeif "}, {"fefef", "refe ", "fewgwerg"},
	}
	for _, strings := range aa {
		fmt.Printf("%v\n",strings)
	}

}
