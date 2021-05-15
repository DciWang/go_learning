package main

import "fmt"

/**
模拟其他语言继承
*/

type animal struct {
	name string
}
type dog struct {
	animal
	feet uint8
}

func (a animal) move() {
	fmt.Printf("%s 会动\n", a.name)
}
func (d dog) wang() {
	fmt.Printf("%s 在:汪汪汪\n", d.name)
}

func main() {
	dd := dog{animal{"q"}, 4}
	//内部拥有的方法，外部也有了
	dd.move() //q 会动
	dd.wang() //q 在:汪汪汪
}
