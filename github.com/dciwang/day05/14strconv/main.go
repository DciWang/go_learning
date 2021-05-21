package main

import (
	"fmt"
	"strconv"
)

/*strconv包实现了基本数据类型与其字符串表示的转换*/
func main() {
	str := "100"

	ii := atoi(str)
	itoa(ii)
}

/*sting  转 int */
func atoi(str string) int {

	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("string convert int failed err : %v \n", err)
	} else {
		fmt.Printf("string convert int success i: %v\n", i)
	}
	return i
}

/*int  转  sting  */
func itoa(i int) {
	str := strconv.Itoa(i)
	fmt.Printf("int convert string ,str: %v\n", str)
	int, _ := strconv.ParseInt(str, 0, 0)
	fmt.Printf("int to string  of   string(int),str: %v\n", string(int))
}

/*
int
int→string
string := strconv.Itoa(int)
1
int→int64
int64_ := int64(int)
1
int64→string
string := strconv.FormatInt(int64,10)
1
int→float
float := float32(int)
float := float64(int)
1
2
int→uint64
uint64 := uint64(int)
1
float
float→string
string := strconv.FormatFloat(float64,'E',-1,64)
string := strconv.FormatFloat(float32,'E',-1,32)
1
2
参数解释：表示格式：‘f’（ddd.dddd）、‘b’（-ddddp±ddd，指数是二进制）、’e’（-d.dddde±dd，指数是十进制）、’E’（-d.ddddE±dd，指数是十进制）、’g’（指数大时,用’e’格式,否则’f’格式）、’G’（指数大时,用’E’格式,否则’f’格式）

float→int64
int64 := int64(float)
1
float→int
int := int(float)
1
string
string→int
int, err := strconv.Atoi(string)
1
string→int64
int64, err := strconv.ParseInt(string, 10, 64)
1
string→float
float,err := strconv.ParseFloat(string,64)
float,err := strconv.ParseFloat(string,32)
1
2
string→bool
bool, err := strconv.ParseBool("true")
1
bool
bool→string
string := strconv.FormatBool(true)
1
interface
interface→int
interface.(int64)
1
interface→string
interface.(string)
1
interface→float
interface.(float64)
interface.(float32)
1
2
interface→bool
interface.(bool)
1
uint
uint64→string
string := strconv.FormatUint(uint64, 10)
*/
