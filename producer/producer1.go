package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)
func main(){

	p2, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer p2.Close()
	go func() {
		for e1 := range p2.Events() {
			switch evs := e1.(type) {
			case *kafka.Message:
				if evs.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", evs.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", evs.TopicPartition)
				}
			}
		}
	}()

	topic1 := "SampleTopic"
	var key string="root12345"

		p2.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic1, Partition: kafka.PartitionAny},
			Value:          []byte(key),
		}, nil)
	

	// Wait for message deliveries before shutting down
	p2.Flush(15 * 1000)


										//SECOND PRODUCER

	p3, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
	panic(err)
	}
	defer p3.Close()
	go func() {
	for e2 := range p3.Events() {
								switch evs1 := e2.(type) {
							case *kafka.Message:
							if evs1.TopicPartition.Error != nil {
								fmt.Printf("Delivery failed: %v\n", evs1.TopicPartition)
									} else {
									fmt.Printf("Delivered message to %v\n", evs1.TopicPartition)
													}
												}
											}
										}()
									
										topic2 := "Topic"
										var val string="1234"
									
											p3.Produce(&kafka.Message{
												TopicPartition: kafka.TopicPartition{Topic: &topic2, Partition: kafka.PartitionAny},
												Value:          []byte(val),
											}, nil)
										
									
										// Wait for message deliveries before shutting down
										p3.Flush(15 * 1000)
									

}