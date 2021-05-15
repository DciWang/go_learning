package main

import (
	"encoding/json"
	"fmt"
)

/**
结构体与json

1.把Go语言中的结构体变量 --->json格式的字符串
2.json格式的字符串  ---->go语言中能够识别的结构体变量
*/
type person struct {
	Name   string `json:"name" db:"name" ini:"name"` //分别对应json，数据库，配置文件
	Age    uint8  `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	wdc := person{
		Name:   "wdc",
		Age:    24,
		Gender: "男",
	}
	b, error := json.Marshal(wdc) //序列化
	if error != nil {
		fmt.Printf("%#v\n", error)
	} else {
		fmt.Println(string(b))
	}

	q := `{"name":"wdc","age":24,"gender":"男"}`
	var p person
	//反序列化
	json.Unmarshal([]byte(q), &p) //传指针是为了能在函数内部修改p的值
	fmt.Println(p)
}
