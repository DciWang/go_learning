package main

import "fmt"

type Mover interface {
	move()
}

type Eater interface {
	eat()
}

type Dog struct {
	name string
}

func (d Dog) move() {
	fmt.Println("狗狗会动")
}

func (d *Dog) eat() {
	fmt.Println("dog can eat")

}

func main() {
	wangcai := &Dog{"wangcai"}
	fugui := Dog{"fugui"}
	var x Mover
	x = wangcai //x can accept  point type  of dog
	x.move()
	x = fugui //x can accept   type  of  dog
	x.move()

	var y Eater
	y = wangcai //y can accept  point type  of dog
	y.eat()
	//y=fugui   can not accept   type  of  dog

}
