package main

import "fmt"

var (
	b1 = false
	b2 = true
	b3 bool //Go中bool默认为false
)

func main() {
	fmt.Printf("b1:%T\n", b1)
	fmt.Printf("b2:%T\n", b2)

	fmt.Println("b3:", b3)
}
