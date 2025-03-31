package go_demo

// 'filtered := numbers[:0]' 创建了一个新的切片 filtered，它与 numbers 共享底层数组，但长度为零，同时保留了 numbers 的容量
// 避免了额外的内存分配，因为我们实际上是修改了 numbers（或者是 numbers 的底层数组）

var originSlice = []int {1,2,3,4,5}


func FilterWithoutAllocation() {
	filtered := originSlice[:0]
	for _, val := range originSlice {
		if val % 2 == 0 {
			filtered = append(filtered, val)
		}
	}
}
