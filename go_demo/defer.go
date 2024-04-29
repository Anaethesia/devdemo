package go_demo

import "fmt"

// defer实现在另一个函数的开头和结尾执行一个函数
func MultiStageDefer() func () {
	fmt.Println("Run init")
	return func() {
		fmt.Println("Run cleanup")
	}
}

func main() {
	defer MultiStageDefer()()
	fmt.Println("Main func called")
}

// 执行顺序 解析MultiStageDefer() 打印Run init，然后拿到return func给defer去执行
// Run init
// Main func called
// Run cleanup