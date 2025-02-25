package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "os/signal"

    "github.com/IBM/sarama"
)

type Message struct {
    Serial    string  `json:"serial"`
    IDPort    int     `json:"idport"`
    Timestamp string  `json:"timestamp"`
    Value     float64 `json:"value"`
}

func main() {
    // Configuração do consumer
    config := sarama.NewConfig()
    config.Consumer.Return.Errors = true
    config.Consumer.Offsets.Initial = sarama.OffsetOldest // Começa a ler desde o início

    // Criação do consumer group
    brokers := []string{"localhost:9092"}
    group := "consumer-01" // Nome do consumer group
    consumer, err := sarama.NewConsumerGroup(brokers, group, config)
    if err != nil {
        log.Fatalf("Erro ao criar o consumer group: %v", err)
    }
    defer func() {
        if err := consumer.Close(); err != nil {
            log.Println("Erro ao fechar o consumer group:", err)
        }
    }()

    // Tópico a ser consumido
    topic := "test-topic"

    // Canal para receber sinais de interrupção
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)

    // Handler para processar mensagens
    handler := &consumerGroupHandler{}

    // Contexto para controle de cancelamento
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Consumo de mensagens
    go func() {
        for {
            err := consumer.Consume(ctx, []string{topic}, handler) // Use o contexto criado
            if err != nil {
                log.Printf("Erro ao consumir mensagens: %v", err)
            }
        }
    }()

    // Espera por um sinal de interrupção
    <-signals
    fmt.Println("Interrompido")
}

// consumerGroupHandler implementa sarama.ConsumerGroupHandler
type consumerGroupHandler struct{}

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
    return nil
}

func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
    return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
    for msg := range claim.Messages() {
        var message Message
        err := json.Unmarshal(msg.Value, &message)
        if err != nil {
            log.Printf("Erro ao deserializar a mensagem JSON: %v", err)
            continue
        }

        log.Printf("Partição: %d, Offset: %d, Mensagem: %+v\n", msg.Partition, msg.Offset, message)

        // Verifica se o serial é "A" ou "B"
        if message.Serial == "A" || message.Serial == "B" {
            log.Printf("Mensagem inválida: %+v\n", message)
            continue
        } else {
			log.Printf("Mensagem válida: %+v\n", message)
	        session.MarkMessage(msg, "")
    	}
    }
    return nil
}