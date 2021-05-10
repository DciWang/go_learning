package main

import (
	"fmt"
	"strings"
)

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".	txt")
	fmt.Println(jpgFunc("aaa"))
	fmt.Println(txtFunc("bbb"))


}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasPrefix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
