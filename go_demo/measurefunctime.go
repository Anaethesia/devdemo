package go_demo

import (
	"fmt"
	"time"
)

// 测量函数的执行时间
func TrackTime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	fmt.Println("elapsed:", elapsed)
	return elapsed
}

func main() {
	defer TrackTime(time.Now())
	time.Sleep(10 * time.Second)
}

