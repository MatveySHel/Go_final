package delivery

import (
	"fmt"
	"github.com/MatveyShel/Go_final/notification/app"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)


func SendToTelegram(kc app.Consumer, message string) error {

	 formattedMessage, err := FormatBookingMessage(message)
		if err != nil {
			return err
		}
	msg := tgbotapi.NewMessage(kc.GetchatID(), formattedMessage)
	if _, err := kc.Getbot().Send(msg); err != nil {
	 return fmt.Errorf("failed to send message to Telegram: %v", err)
	}
	return nil
   }
   