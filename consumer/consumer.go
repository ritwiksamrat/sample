package main

import (
	
	"fmt"
	// "log"
	// "os"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"github.com/ritwiksamrat/consumer/consumer1"
)

func main() {
	
	fmt.Println(consumer1.samp)
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "smallest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"SampleTopic"}, nil)
	for {
		key,err:= c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", key.TopicPartition, string(key.Value))
			
		}else {
			fmt.Printf("Consumer error: %v (%v)\n", err, key)
		}
	}

defer c.Close()

}


