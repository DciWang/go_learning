package main

import "fmt"

func main() {
ff1(ff3(ff2,100,200))

}
func ff1(f func())  {
	fmt.Println("this is ff1")
	f()
}

func ff2(x,y int)  {
	fmt.Println("this is ff2")
	fmt.Println("x+y=",x+y)
}

/*需求：
把ff2当参数传到ff1中
*/

func ff3(f func(int,int),x,y int) func(){
	return func() {
		f(x,y)
	}
}

func f1(f func())  {
	fmt.Println("this is f1")
	f()
}

