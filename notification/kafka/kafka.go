package kafka

import (
 "context"
 "fmt"
 "log"

 tgbotapi "gopkg.in/telegram-bot-api.v4"
 "github.com/MatveyShel/Go_final/notification/delivery"
)


func (kc *KafkaConsumer) GetchatID() int64{
    return kc.chatID
}

func (kc *KafkaConsumer) Getbot() *tgbotapi.BotAPI{
    return kc.bot
}

func (kc *KafkaConsumer) StartConsuming(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            log.Println("Завершение работы StartConsuming по сигналу контекста.")
            return
        default:

            msg, err := kc.reader.FetchMessage(ctx)
            if err != nil {
                log.Printf("Ошибка чтения сообщения из Kafka: %v", err)
                continue
            }

            fmt.Printf("Получено сообщение: %s\n", string(msg.Value))

            err = delivery.SendToTelegram(kc, string(msg.Value))
            if err != nil {
                log.Printf("Ошибка при отправке сообщения в Telegram: %v", err)
                continue
            }

            if err := kc.reader.CommitMessages(ctx, msg); err != nil {
                log.Printf("Ошибка при коммите сообщения: %v", err)
            } else {
                log.Println("Сообщение успешно отправлено в Telegram и закоммичено.")
            }
        }
    }
}

func (kc *KafkaConsumer) Close() error {
	if err := kc.reader.Close(); err != nil {
        return err
	  }
    return nil
}