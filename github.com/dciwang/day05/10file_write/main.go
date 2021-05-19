package main

import (
	"fmt"
	"os"
)

func main() {
	fileWrite()
}

func fileWrite() {
	file, err := os.Create("./newFile")
	if err != nil {
		fmt.Printf("create file failed, err:%v\n", err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("hello  world \n")
		file.Write([]byte("hello dc \n"))
	}

}
