package main

import "fmt"

type dog struct {
	name string
}

type eat interface {
	eat()
}
type mover interface {
	move()
}

//nested（嵌套）  tow  interface
type doger interface {
	eat
	mover
}

func (d dog) move() {
	fmt.Println("dog is  moving")
}
func (d dog) eat() {
	fmt.Println("dog is eating")
}

func main() {
	wangcai := dog{
		"wangcai",
	}
	var doger doger
	doger = wangcai
	doger.eat()
	doger.move()
}
