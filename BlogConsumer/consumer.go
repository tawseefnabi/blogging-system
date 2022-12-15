package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	topic               string
	conn                sarama.Consumer
	BlogConsumerService *BlogConsumerService
}

func NewConsumer(topic string, broker []string, BlogConsumerService *BlogConsumerService) *Consumer {
	conn, err := connectConsumer(broker)

	if err != nil {
		log.Panicln(err)
	}
	return &Consumer{
		topic:               topic,
		conn:                conn,
		BlogConsumerService: BlogConsumerService,
	}
}

func (c *Consumer) Init() {
	worker := c.conn
	// Calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(c.topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer Started")

	sigchan := make(chan os.Signal, 1)

	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgcount := 0

	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgcount, string(msg.Topic), string(msg.Value))
				messages := msg.Value
				var payload BlogPayload
				err = json.Unmarshal(messages, &payload)
				if err != nil {
					log.Println("Caught error: ", err.Error())
				} else {
					if payload.RequestType == "create" {
						err := c.BlogConsumerService.WriteData(payload.Body, payload.CallBackUrl)
						if err != nil {
							log.Println("Caught error on indexing: ", err.Error())
						} else {
							log.Println("Data Ingested")
						}
					} else if payload.RequestType == "delete" {
						err := c.BlogConsumerService.DeleteData(payload.Body, payload.CallBackUrl)
						if err != nil {
							log.Println("Caught error in deleting: ", err.Error())
						} else {
							log.Println("Data Removed")
						}
					}
				}
			case <-sigchan:
				fmt.Println("Interrupt in detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgcount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}
}
