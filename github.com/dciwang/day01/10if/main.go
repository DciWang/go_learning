package main

import "fmt"

func main() {

	//age 作用域为 main函数
		age := 19
	if age>18 {
		fmt.Println("澳门首家线上赌场上线")
	}
	fmt.Println("该写暑假作业了")
	// age1作用域为if语句
	if age1 := 18;age1>35  {
		fmt.Println("人到中年")
	}else if age >18 {
		fmt.Println("人到青年")
	}else {
		fmt.Println("好好学习")
	}


	


}
