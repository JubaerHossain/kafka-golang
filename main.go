package main

import (
	"fmt"
	"kafkago/kafka"
	"time"
)

func main() {
	fmt.Println("Okay...")
	kafka.StartKafka()

	fmt.Println("Kafka has been started...")

	time.Sleep(10 * time.Minute)
}
