package main

import "fmt"

func main() {
	var (
		topic  string   = "blogs"
		broker []string = []string{"localhost:9092"}
	)

	fmt.Println("Consumer Started")
	BlogConsumerService := BlogConsumerService{}
	consumer := NewConsumer(topic, broker, &BlogConsumerService)
	consumer.Init()

}
