package main

import "fmt"
/*
Go语言中函数的return不是院子操作，在底层粉我两部来执行
第一步：返回值赋值
第二部：真正的RET返回

函数中如果存在defer，那么defer执行的时机是在第一步跟第二部
*/
func f1() int {
	x := 5  //此时的x已经为5
	defer func() {
		x++
	}()
	return x    //将5赋值给返回值，然后再对x+=
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5   //将5赋值给x，然后对x++，最后返回，得到6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x    //将x赋值给y，然后对x++，但是返回y
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)     //x为外面函数x的一个拷贝值
	return 5     //将5赋值给x，然后返回
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}