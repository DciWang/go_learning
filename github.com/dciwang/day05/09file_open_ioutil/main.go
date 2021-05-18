package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	readFileByIoutil()
}

func readFileByIoutil() {
	file, err := ioutil.ReadFile("./main.go")

	if err == io.EOF {
		return
	}

	if err != nil {
		fmt.Printf("file %v open faild \n", err)
		return
	}
	fmt.Println(string(file))

}
