package main

import (
	"fmt"
	"sort"
)

func main() {

	var m map[string]string    //定义但是没有初始化(没有在内存中开辟空间)
	m1 := map[string]int{}

	fmt.Println(m == nil)
	fmt.Println(m1 == nil)


	fmt.Println(m1)
	m2 :=make(map[string]int,10)   //估算好该map的容量，避免程序运行的时候进行洞体啊扩容
	m2["dciwang"]=26
	m2["Q"]=25
	/*判断是否存在*/
	value,ok :=m2["曹家"]
	if !ok {
		fmt.Println("查无此人")
	}else {
		fmt.Println(value)
	}
	/*遍历*/
	//遍历键值对
	for k, v := range m2 {
		fmt.Println(k,v)
	}
	//遍历键
	for k, _ := range m2 {
		fmt.Println(k)
	}
	//遍历值
	for _, v := range m2 {
		fmt.Println(v)
	}
	/*删除*/
	delete(m2, "dciwang")
	delete(m2, "Caojia")   //如果没有不存在的键值对，什么都不干

	/*按照指定顺序遍历key*/
	m3 := make(map[int]string,16)
	m3[3]="aaa"
	m3[2]="aaaa"
	m3[4]="aaaaa"
	m3[6]="www"
	m3[8]="aaaaaa"
	m3[9]="ggg"
	//遍历 map 的key添加到切片中
	i := []int{}
	for k := range m3 {
		i = append(i, k)
	}
	//对切片进行排序
	sort.Ints(i)
	//根据key你到value
	for _, v := range i {
		fmt.Printf("m3_v:%d,m3_v:%s\n",v,m3[v])
	}


	/*元素为map类型的切片*/
	s1 := make([]map[string]int,10,10)    //初始化切片但是没有初始化map

	s1[0] = map[string]int{}
	s1[0]["a"] =2
	s1[0]["b"] =3
	s1[1] = map[string]int{}
	s1[1]["b"] =3

	fmt.Println(s1)
	/*值为切片类型的map*/


	m4 :=make(map[string][]int,10)

	m4["aa"] = []int{111}
	fmt.Println(m4)



	mm :=make(map[string][]string,10)

	for k, _ := range mm {
		mm[k]=make([]string,10,10)
	}
	mm["aaa"]=[]string{"qqq","sss"}
	mm["bb"]=[]string{"qqq","sss"}

	fmt.Println(mm)

}
