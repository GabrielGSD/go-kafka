package main

import (
    "encoding/json"
    "math/rand"
    "log"
    "sync"
    "time"

    "github.com/IBM/sarama"
)

type Message struct {
    Serial    string  `json:"serial"`
    IDPort    int     `json:"idport"`
    Timestamp string  `json:"timestamp"`
    Value     float64 `json:"value"`
}

func publishMsg(producer sarama.SyncProducer, serial []string, i int) {
    message := Message{
        Serial:    serial[rand.Intn(len(serial))],
        IDPort:    1,
        Timestamp: time.Now().Format("2006-01-02 15:04:05"),
        Value:     rand.Float64() * 100,
    }
    value, err := json.Marshal(message)
    if err != nil {
        log.Fatalf("Erro ao serializar a mensagem: %v", err)
    }

    msg := &sarama.ProducerMessage{
        Topic: "test-topic",
        Value: sarama.StringEncoder(value),
    }

    partition, offset, err := producer.SendMessage(msg)
    if err != nil {
        log.Printf("Erro ao enviar mensagem: %v", err)
    } else {
        log.Printf("G%d - Mensagem enviada para a partição %d com offset %d", i, partition, offset)
    }
}

func main() {
    // Configuração do produtor Kafka
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true

    // Criação do produtor
    producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
    if err != nil {
        log.Fatalf("Erro ao criar o produtor: %v", err)
    }
    defer producer.Close()

    serial := []string{"A", "B", "C", "D", "E"}

    numGoroutines := 10
    var wg sync.WaitGroup
    wg.Add(numGoroutines)

    for i := 0; i < numGoroutines; i++ {
        go func(i int) {
            defer wg.Done()
            for {
                publishMsg(producer, serial, i)
                time.Sleep(2 * time.Second)
            }
        }(i)
    }

    wg.Wait()
}