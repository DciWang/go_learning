package main

import "fmt"

//for  循环
func main() {
	//基本语句

	s := "hello 多聪"
	//j :=0
	for i := 0; i < len(s); i++ {
		j := 0
		for ; j < len(s); {
			fmt.Printf("%c\n", s[j])
			j++
			fmt.Println("m=", i)
			fmt.Println("j=", j)
		}
	}

	//死循环
	/*for {
		fmt.Println("死循环了")
	}*/

	/*
		for  range循环， 用来遍历字符串、数组、切片等
	*/
	fmt.Println("----------------------")
	for i, i2 := range s {
		fmt.Printf("m=%d i2=%c\n", i, i2)
	}
	fmt.Println("----------------------")
	/*跳出for循环*/

	//1. 信号量判断跳出for循环
	var semaphore bool

	for i := 1; i < 10; i++ {
		for j := 1; j < j; j++ {
			if j ==2 {
				semaphore = true
				break
			}
			fmt.Printf("m=%d,j=%d,", i, j)
		}
		if semaphore {
			break
		}
	}

	//2. goto跳出for循环

	for m := 0; m <10 ; m++ {
		for k := 0; k < m; k++ {
			if k ==2 {
				goto  other
			}
			fmt.Printf("m=%d,k=%d,", m, k)
		}
	}
	other:
		fmt.Println("goto跳出循环")

	//3.continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。 在 continue语句后添加标签时，表示开始标签对应的循环

forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

}
