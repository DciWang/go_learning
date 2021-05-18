package main

import (
	"fmt"
	"io"
	"os"
)

/*打开一个文件 os 包*/
func main() {

	FileObj, err := os.Open("./main.go")
	defer FileObj.Close() //记得关闭文件
	if err != nil {
		fmt.Printf("open file filed error:%v\n", err)
		return
	}
	//读文件
	tmp := make([]byte, 128)
	for {
		n, error := FileObj.Read(tmp)
		if error == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if error != nil {
			fmt.Printf("read  file filed error:%v\n", error)
		}
		fmt.Printf("读到了%d个字节\n", n)
		fmt.Printf("this is value of FileObj : %#v\n", tmp[:n])
		if n < 128 {
			return
		}
	}
}
