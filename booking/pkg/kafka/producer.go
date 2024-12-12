package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:      brokers,
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
		}),
	}
}

func (p *Producer) SendMessage(ctx context.Context, value []byte) error {
	msg := kafka.Message{
		Value: value,
	}
	err := p.writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Printf("fail to send message into Kafka: %v", err)
		return fmt.Errorf("fail to send message into Kafka: %w", err)
	}
	log.Printf("Message sent to Kafka: %s", value)
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}