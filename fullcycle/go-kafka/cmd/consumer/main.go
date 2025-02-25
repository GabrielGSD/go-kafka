package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "go-kafka-kafka-1:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest",
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Println("error consumer", err.Error())
	}
	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil) // posso consumir mais de um tópico
	for {
		msg, err := c.ReadMessage(-1) // -1 = tempo de espera infinito
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
