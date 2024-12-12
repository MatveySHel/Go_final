package kafka


import (
	"github.com/segmentio/kafka-go"
	"context"
	"log"
	"fmt"
)

   
type KafkaConsumer struct {
	reader              *kafka.Reader
}

// Инициализация KafkaConsumer с конфигурацией и хэндлером
func NewKafkaConsumer(broker, topic, groupID string) (*KafkaConsumer, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupID,
	})

	return &KafkaConsumer{
		reader: r,
	}, nil
}

// Запуск процесса чтения сообщений из Kafka
func (kc *KafkaConsumer) StartConsuming(ctx context.Context) {
	for {
		msg, err := kc.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message from Kafka: %v", err)
			continue
		}

		// Передаём сообщение в обработчик уведомлений
		fmt.Printf("Получено сообщение: %s\n", string(msg.Value))
	}
}