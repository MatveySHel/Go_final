package main

import (
	"log"
	"os"

	tbm "github.com/MatveyShel/Go_final/api/telegram_bot"
	"github.com/joho/godotenv"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}
	hotelService := os.Getenv("HOTEL_SERVICE_ADDR")
	bookingService := os.Getenv("BOOKING_SERVICE_ADDR")
	token := os.Getenv("TELEGRAM_BOT_TOKEN_HANDLE")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal("Failed to get updates channel:", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		userMessage := update.Message.Text

		if state, ok := tbm.UserStates[chatID]; ok {
			switch state.Step {
			case 1, 2, 3, 4:
				tbm.HandleBookState(bot, chatID, userMessage, state, bookingService)
			case 5:
				tbm.HandleBookingsState(bot, chatID, userMessage, state, bookingService)
			default:
				bot.Send(tgbotapi.NewMessage(chatID, "Неизвестное состояние. Попробуйте заново."))
				delete(tbm.UserStates, chatID)
			}
			continue
		}


		switch userMessage {
		case "/start":
			tbm.HandleStartCommand(bot, chatID)
		case "/hotels":
			tbm.HandleHotelsCommand(bot, chatID, hotelService)
		case "/book":
			tbm.HandleBookCommand(bot, chatID)
		case "/bookings":
			tbm.HandleBookingsCommand(bot, chatID)
		default:
			bot.Send(tgbotapi.NewMessage(chatID, "Неизвестная команда. Попробуйте /start"))
		}
	}
}
