package main

import "fmt"

/**
递归
一定要有一个明确的退出条件
递归适用于那种问题相同\问题的规模越来越小的问题

 */
func main() {

   fmt.Println( f(6))
   fmt.Println( ff(3))
}



/**
阶乘
 */

func f(n uint64) uint64{
	if n<=1 {
		return 1
	}
	return n * f(n-1)
}


/*上台阶问题：
	n个台阶，一次可以走一步，也可以走两步，有多少种走法
*/

func ff(n uint64) uint64  {
	if n==1 {
		return  1
	}
	if n==2 {
		return 2
	}

	return ff(n-1)+ff(n-2)
}