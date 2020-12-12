package main

import (
	
	"fmt"
	// "log"
	// "os"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var samp string

func main() {


	d, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "smallest",
	})

	if err != nil {
		panic(err)
	}
	d.SubscribeTopics([]string{"Topic"}, nil)
	for {
		key,err:= d.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", key.TopicPartition, string(key.Value))
			samp=string(key.Value)
			
		}else {
			fmt.Printf("Consumer error: %v (%v)\n", err, key)
		}
	}
defer d.Close()
}


