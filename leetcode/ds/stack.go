package ds


type Stack []int

func(s *Stack) Push(x int) {
	*s = append(*s, x)
}

func(s *Stack) Pop() int {
	length := len(*s) - 1
	x := (*s)[length]
	*s = (*s)[:length]
	return x
}

func(s *Stack) Top() int {
	length := len(*s) - 1
	return (*s)[length]
}

func(s *Stack) Empty() bool {
	return len(*s) == 0
}

type Queue []int

func(q *Queue) LPush(x int) {
	*q = append([]int{x}, *q...)
}

func(q *Queue) RPush(x int) {
	*q = append(*q, x)
}

func(q *Queue) LPop() int {
	if len(*q) == 0 {
		return 0
	}
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func(q *Queue) RPop() int {
	if len(*q) == 0 {
		return 0
	}
	length := len(*q) - 1
	x := (*q)[length]
	*q = (*q)[:length]
	return x
}

func(q *Queue) Front() int {
	if len(*q) == 0 {
		return 0
	}
	return (*q)[0]
}

func(q *Queue) Tail() int {
	if len(*q) == 0 {
		return 0
	}
	return (*q)[len(*q)-1]
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}