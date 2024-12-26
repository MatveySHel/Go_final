package main

import (
  "context"
  "fmt"
  "log"
  "os"
  "os/signal"
  "syscall"
  "strconv"

  "github.com/joho/godotenv"
  "github.com/MatveyShel/Go_final/notification/kafka"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  broker := os.Getenv("BROKER_ADDR")
  topic := "notifications"
  groupID := "notification-group"
  botToken := os.Getenv("TELEGRAM_BOT_TOKEN_NOTIFICATIONS")
  chat := os.Getenv("TELEGRAM_CHAT")
  chatID, err  := strconv.ParseInt(chat, 10, 64) 
  if err!=nil {
    log.Println("Ошибка преобразования:", err)
  }

  consumer, err := kafka.NewKafkaConsumer(broker, topic, groupID, botToken, chatID)
  if err != nil {
    log.Fatalf("Не удалось создать KafkaConsumer: %v", err)
  }
  fmt.Println("KafkaConsumer успешно создан!")

  ctx, cancel := context.WithCancel(context.Background())
  done := make(chan bool)

  go func() {
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
    <-stop
    fmt.Println("Получен сигнал завершения работы, остановка KafkaConsumer...")
    cancel()
    done <- true
  }()

  go func() {
    consumer.StartConsuming(ctx)
  }()

  <-done
  fmt.Println("Завершаем приложение...")

  
  if err := consumer.Close(); err != nil {
    log.Printf("Ошибка при закрытии KafkaReader: %v", err)
  }

  fmt.Println("Приложение завершено.")
}
