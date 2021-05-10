package main

import "fmt"

/**
1. uint8类型，或者叫byte类型。代表了一个ASCII的一个人字符
2. rune类型，代表的了一个utf-8字符，用来处理中文，韩文等，实际上是一个int32
 */
func main() {
	s := "hello China"

	for i :=0; i<len(s);i++{
		fmt.Println(s[i])
		fmt.Printf("%c\n",s[i])

	}

	for _, i2 := range s {
		fmt.Printf("%c\n",i2)
	}

	/**
	修改字符串：
		字符串本身是不能被修改的，如果非要修改，要强制转换成[]rune或者[]byte数组，完后再转换成string，无论哪种转换都会重新分配内存，并复制字节数组
	 */

	//转成[]byte数组
	byte1 := []byte(s)
	byte1[0]='H'
	fmt.Println(string(byte1))

	//转成[]rune数组

	rune1 :=[]rune(s)
	rune1[1]='E'
	fmt.Println(string(rune1))
}
