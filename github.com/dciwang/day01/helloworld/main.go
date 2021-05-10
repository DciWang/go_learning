//包，如果是 main 包，则表明是一个可执行包，否则不是
package  main
//导入语句
import "fmt"

//函数外只能对（变量、常量、函数）的声明

// main函数，程序的入口函数
func main() {
fmt.Println("hello world")

 FmtPrint()
}

func FmtPrint()  {
	fmt.Println("人生苦短，Let's Go ！！！")
}


