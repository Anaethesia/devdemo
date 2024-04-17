package main

import (
	"devdemo/go_demo"
)

func main() {
	// kafka_demo.Kafka_main()

	go_demo.ProduceJob()


	go_demo.ConsumeJob()


}