package go_demo

import (
	"fmt"
	"unsafe"
)

// uintptr、unsafe.Pointer和普通指针之间的转换关系
// uintptr <==> unsafe.Pointer <==> *T

// unsafe pointer 获取结构体成员
func main() {

	var x struct {
		a int
		b int
		c []int
	}

	// unsafe.Offsetof 函数的参数必须是一个字段,  比如 x.b,  方法会返回 b 字段相对于 x 起始地址的偏移量, 包括可能的空洞。

	// 指针运算 uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)。

	// 和 pb := &x.b 等价
	pb := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))

	*pb = 42
	fmt.Println(x.b) // "42"
}