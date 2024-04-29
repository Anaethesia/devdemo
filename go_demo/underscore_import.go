package go_demo

import "fmt"

// import (
//	_ "pkg/github.com/xxx"
//)

// import 包加下划线的话， 在不创建包引用的情况下，执行包里面的初始化init函数

func init() {
	fmt.Println("init call")
}