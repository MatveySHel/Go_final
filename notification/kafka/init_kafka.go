package kafka

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	
	tgbotapi "gopkg.in/telegram-bot-api.v4"
   )
   
   type KafkaConsumer struct {
	reader    *kafka.Reader
	bot       *tgbotapi.BotAPI
	chatID    int64 // ID пользователя или чата в Telegram
   }
   
   // Инициализация KafkaConsumer с конфигурацией и хэндлером
   func NewKafkaConsumer(broker, topic, groupID, botToken string, chatID int64) (*KafkaConsumer, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
	 Brokers: []string{broker},
	 Topic:   topic,
	 GroupID: groupID,
	})
   
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
	 return nil, fmt.Errorf("failed to create telegram bot: %v", err)
	}
   
	return &KafkaConsumer{
	 reader: r,
	 bot:    bot,
	 chatID: chatID,
	}, nil
   }