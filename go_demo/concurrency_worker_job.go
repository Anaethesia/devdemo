package go_demo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// From apple boy's go training course
// Sample code about how to use worker and jobs for concurrency coding

var concurrencyProcess = 10
var jobCount = 100

func ConcurrencyWorkerJob() {
	var wg sync.WaitGroup
	wg.Add(jobCount)

	found := make(chan int)
	queue := make(chan int)

	// 无缓冲队列，预计导入100个job
	go func(queue chan<- int) {
		for i := 0; i < jobCount; i++ {
			queue <- i
		}
		close(queue)
	}(queue)

	// 开10个goroutine，每个go routine都在读queue channel拿值执行逻辑，然后结果塞到found channel中
	for i := 0; i < concurrencyProcess; i++ {
		go func(queue <-chan int) {
			for val := range queue {
				defer wg.Done()

				waitTime := rand.Int31n(1000)
				fmt.Println("job: ", val, "wait time: ", waitTime, "ms")
				time.Sleep(time.Duration(waitTime) * time.Millisecond)
				found <- val
			}

		}(queue)
	}

	go func() {
		wg.Wait()
		close(found)
	}()

	var results []int

	// found channel一有数据就输出并导结果到results的slice中
	for f := range found {
		fmt.Println("finished job:", f)
		results = append(results, f)
	}

	fmt.Println("results: ", results)
}
