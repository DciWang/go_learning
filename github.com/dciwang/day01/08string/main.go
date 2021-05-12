package main

import (
	"fmt"
	"strings"
)

/*
Go语言里的字符串内部实现用utf-8编码，只能是  ""  包裹起来的！！！， '' 包裹起来的是字符


*/
func main() {
	s1 := "D:\\dciwang\\src"
	/*
		字符串转义
		\n   换行
		\r   回车
		\t	 制表符
		\'	 单引号
		\"	 双引号
		\\	 反斜杠

	*/

	fmt.Printf("s: %s\n", s1)
	//多行字符串
	s2 := `
	世情薄
	人情恶
	雨
`
	fmt.Printf(s2)
	ss1 := "DciWang"
	ss2 := "大帅比"
	//字符串操作
	fmt.Println(len(s2))         //字符串长度
	fmt.Printf("%s%s", ss1, ss2) //不做拼接
	fmt.Println()
	ss := fmt.Sprintf("%s%s", ss1, ss2) //只做拼接但不打印
	fmt.Println(ss)
	fmt.Println()
	fmt.Println(ss1 + ss2) //拼接并打印

	//分割
	s3 := `D:dciwang\src\main`
	s4 := strings.Split(s3, "\\")
	fmt.Println(s4)
	//判断是否包含
	s5 := strings.Contains(s3, "Q")
	fmt.Println(s5)
	//判断是否为 D  开头
	fmt.Println(strings.HasPrefix(s3, "D"))
	//判断是否为 D  结尾
	fmt.Println(strings.HasSuffix(s3, "D"))
	fmt.Println(s5)

	//判断字串第一次出现的位置
	fmt.Printf("a第一次出现的位置为: %v\n", strings.Index(s3, "a"))
	//如果没有出现则返回-1
	fmt.Printf("b第一次出现的位置为: %v\n", strings.Index(s3, "b"))
	//判断字串最后一次出现的位置
	fmt.Printf("a最后一次出现的位置为：%v\n", strings.LastIndex(s3, "a"))
	//拼接
	fmt.Println(strings.Join(s4, "+"))

}
