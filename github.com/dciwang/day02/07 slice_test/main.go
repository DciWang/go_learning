package main

import (
	"fmt"
	"sort"
)

/*练习*/
func main() {

	s := make([]int,5,10)

	for i := 0; i <10 ; i++ {
		s = append(s,i)
	}

	a :=[...]int{1,2,5,4,3,6}
	fmt.Println(s)
	sort.Ints(a[:])
	fmt.Println(a)
}
