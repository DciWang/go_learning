package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin(distribution)
	fmt.Println("剩下：", left)
	fmt.Println("distribution：", distribution)
}

func dispatchCoin(distribution map[string]int) (left int) {
	for _, user := range users {
		for _, i := range user {
			switch i {
			case 'e', 'E':
				distribution[user]++
				coins--
			case 'i', 'I':
				distribution[user] += 2
				coins -= 2
			case 'o', 'O':
				distribution[user] += 3
				coins -= 3
			case 'u', 'U':
				distribution[user] += 4
				coins -= 4
			}
		}
	}
	left = coins
	return
}
