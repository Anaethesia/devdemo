package go_demo

import (
	"fmt"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

func ProduceJob() {
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		fmt.Println("Produce job err: ", err)
	}

	id, err := c.Put([]byte("hello"), 0, 0, 120*time.Second)
	if err != nil {
		fmt.Println("the err is", err)
		return
	}
	fmt.Println("the produced job id is ", id)
	id1, err := c.Put([]byte("hello1"), 1, 0, 120*time.Second)
	if err != nil {
		fmt.Println("the err is", err)
		return
	}
	fmt.Println("the produced job id is ", id1)
	id2, err := c.Put([]byte("hello2"), 2, 0, 120*time.Second)
	if err != nil {
		fmt.Println("the err is", err)
		return
	}
	fmt.Println("the produced job id is ", id2)
	id3, err := c.Put([]byte("hello3"), 3, 0, 120*time.Second)
	if err != nil {
		fmt.Println("the err is", err)
		return
	}
	fmt.Println("the produced job id is ", id3)
	id4, err := c.Put([]byte("hello4"), 4, 0, 120*time.Second)
	if err != nil {
		fmt.Println("the err is", err)
		return
	}
	fmt.Println("the produced job id is ", id4)
}

func ConsumeJob() {
	c, _ := beanstalk.Dial("tcp", "127.0.0.1:11300")
	for {
		id, body, err := c.Reserve(30 * time.Second)
		if err != nil {
			fmt.Println("the err is", err)
			return
		}
		fmt.Println("the consumed job id is ", id)
		fmt.Println("the consumed body is ", string(body))
	}

}