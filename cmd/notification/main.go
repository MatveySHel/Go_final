package main

// Инициализация и запуск микросервиса

import (
	"github.com/MatveyShel/Go_final/notification/kafka"
	"context"
	"log"
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	broker := os.Getenv("BROKER_ADDR")
	topic := "notifications"   // Топик для чтения
	groupID := "notification-group" // Группа потребителей
   
	consumer, err := kafka.NewKafkaConsumer(broker, topic, groupID)
	if err != nil {
	 log.Fatalf("Не удалось создать KafkaConsumer: %v", err)
	}
	fmt.Println("KafkaConsumer успешно создан!")
   
	ctx, cancel := context.WithCancel(context.Background())
   
	// Обработка системных сигналов (для корректной остановки)
	go func() {
	 stop := make(chan os.Signal, 1)
	 signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	 <-stop
	 fmt.Println("Получен сигнал завершения работы, остановка KafkaConsumer...")
	 cancel()
	}()
   
	// Запускаем чтение сообщений из Kafka
	go func() {
	 consumer.StartConsuming(ctx)
	}()

	select {
	case <-ctx.Done():
	 fmt.Println("Завершаем приложение...")
	}
   
	time.Sleep(2 * time.Second) // Небольшая пауза
	fmt.Println("Приложение завершено.")
   }