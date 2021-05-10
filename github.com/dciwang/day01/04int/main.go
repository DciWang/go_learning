package main

import "fmt"


/*
一般整型
    int8  带符号的8位整形
    int16  带符号的16位整形
    int32  带符号的32位整形
    int64  带符号的64位整形

    uint8    不带符号的8位整形
    uint16   不带符号的16位整形
    uint32   不带符号的32位整形
    uint64   不带符号的64位整形
特殊整形：
	int	    跟随操作系统而定，32位操作系统位int32，64位操作系统位int64
	uint	跟随操作系统而定，32位操作系统位uint32，64位操作系统位uint64
	uintptr  无符号整型，用于存放一个指针
*/

var (
	value1 = 100  //十进制   默认Go语言中整型位int类型
	value2 = 0100  // 八进制
	value3 = 0x100  // 十六进制
)
func main() {

	fmt.Printf("value1:%d\n", value1)
	fmt.Printf("value2:%d\n", value2)
	fmt.Printf("value3:%d\n", value3)
	var  uInt uint  =88
	var  uInt8 uint  =88
	  uInt   =888
	 fmt.Printf( "%d\n",uInt)
	 fmt.Printf( "%d\n",uInt8)

	  i1 := int8(101)     //明确指定为int8类型，否则默认为int类型

	  /*
	  	%d  转换成10进制
	  	%b  转换成2进制
	  	%o  转换成8进制
	  	%x  转换成16进制
	  */
	  fmt.Printf("i1:%d\n",i1)
}
