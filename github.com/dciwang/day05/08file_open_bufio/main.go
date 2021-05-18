package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	readFormFileByBufio()

}

/*利用bufio这个包读文件*/
func readFormFileByBufio() {
	file, err := os.Open("./main.go")
	defer file.Close()
	if err != nil {
		fmt.Printf("file %v open faild \n", err)
	}
	reader := bufio.NewReader(file)
	for {

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed %v\n", err)
			return
		}
		if err != nil {
			fmt.Printf("file %v open faild \n", err)
		}
	}
}
