package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writeBybufio()
}

func writeBybufio() {
	file, err := os.Create("./writeBybufio")
	if err != nil {
		fmt.Printf("create file failed,err:\n", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("地铁上男人女人面对背贴在一起，你都不知道他们是真情侣关系还是假的。他们不说一句话，但是不得不贴的那么近。前天老大叔小心翼翼拍的那棵杏树，我今天抬头看见一个透漏着黄的果实 \n")
	}
	writer.Flush()
}
