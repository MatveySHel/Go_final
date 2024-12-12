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
)

func main() {
	// Конфигурация Kafka
	broker := "0.0.0.0:9092" // Адрес Kafka брокера
	topic := "notifications"   // Топик для чтения
	groupID := "notification-group" // Группа потребителей
   
	// Создаем Kafka consumer
	consumer, err := kafka.NewKafkaConsumer(broker, topic, groupID)
	if err != nil {
	 log.Fatalf("Не удалось создать KafkaConsumer: %v", err)
	}
	fmt.Println("KafkaConsumer успешно создан!")
   
	// Создаем контекст с отменой для корректного завершения работы consumer'а
	ctx, cancel := context.WithCancel(context.Background())
   
	// Обработка системных сигналов (для корректной остановки)
	go func() {
	 stop := make(chan os.Signal, 1)
	 signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM) // Ловим SIGINT и SIGTERM
	 <-stop
	 fmt.Println("Получен сигнал завершения работы, остановка KafkaConsumer...")
	 cancel() // Завершаем контекст
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