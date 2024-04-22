package go_demo

import "fmt"

// for range 中的 val对应的变量地址都是一样的，每次循环都是把值复制给val地址中，开发中不要记录每次循环的val地址，你会发现都是一样的。

/*
	the address of val is  0xc00001e098
	the value of val is 1
	the address of val is  0xc00001e098
	the value of val is 2
	the address of val is  0xc00001e098
	the value of val is 3
	...
*/
func AddressInSliceMember() {
	temp := []int{1, 2, 3, 4, 5}
	for _, val := range temp {
		fmt.Println("the address of val is ", &val)
		fmt.Println("the value of val is", val)
	}
}
