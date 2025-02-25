package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("transferiu", "teste", producer, []byte("transferencia2"), deliveryChan)
	go DeliveryReport(deliveryChan) //async
	producer.Flush(5000)
	// // é sincrono, então só vai passar para o próximo passo quando a mensagem for entregue
	// e := <-deliveryChan
	// msg := e.(*kafka.Message)
	// if msg.TopicPartition.Error != nil {
	// 	log.Println("Erro ao enviar a mensagem", msg.TopicPartition.Error)
	// } else {
	// 	log.Println("Mensagem enviada", msg.TopicPartition)
	// }

	// producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "go-kafka-kafka-1:9092",
		"delivery.timeout.ms": 0,
		"acks":                "all", // 1 = só confirma que a mensagem foi recebida pelo broker, 0 = não confirma nada, all = confirma que a mensagem foi recebida por todos os brokers
		"enable.idempotence":  true,  // garante que a mensagem seja enviada apenas uma vez e na ordem correta (só funciona com acks=all), mas pode causar lentidão.
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

// O value (msg) é um array de bytes, então não necessariamente precisa ser uma string, pode ser um json, por exemplo.
func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Println("Erro ao enviar a mensagem", ev.TopicPartition.Error)
			} else {
				log.Println("Mensagem enviada", ev.TopicPartition)
				// anotar no banco de dados que a mensagem foi processada
				// ex: confirma que uma transferencia bancária foi realizada
			}
		}
	}
}
