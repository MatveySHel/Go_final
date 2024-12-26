package telegram_bot

import (
	"fmt"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func HandleStartCommand(bot *tgbotapi.BotAPI, chatID int64) {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/hotels"),
			tgbotapi.NewKeyboardButton("/book"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/bookings"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Добро пожаловать! Выберите команду:")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}


func HandleHotelsCommand(bot *tgbotapi.BotAPI, chatID int64, hotelService string) {
	hotels, err := GetHotels(hotelService)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "Ошибка при получении списка отелей"))
		return
	}
	bot.Send(tgbotapi.NewMessage(chatID, "Список отелей:\n"+hotels))
}


func HandleBookCommand(bot *tgbotapi.BotAPI, chatID int64) {
	bot.Send(tgbotapi.NewMessage(chatID, "Введите имя клиента:"))
	UserStates[chatID] = &BookingState{
		Step:    1,
		Booking: Booking{},
	}
}


func HandleBookState(bot *tgbotapi.BotAPI, chatID int64, userMessage string, state *BookingState, bookingService string) {
	switch state.Step {
	case 1:
		state.Booking.Client = userMessage
		state.Step = 2
		bot.Send(tgbotapi.NewMessage(chatID, "Введите название отеля:"))
	case 2:
		state.Booking.Hotel = userMessage
		state.Step = 3
		bot.Send(tgbotapi.NewMessage(chatID, "Введите дату заезда (в формате YYYY-MM-DD):"))
	case 3:
		state.Booking.CheckIn = userMessage
		state.Step = 4
		bot.Send(tgbotapi.NewMessage(chatID, "Введите дату выезда (в формате YYYY-MM-DD):"))
	case 4:
		state.Booking.CheckOut = userMessage
		_, err := BookHotel(bookingService, state.Booking)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, "Ошибка при бронировании"))
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, "Бронирование выполнено успешно!"))
		}
		delete(UserStates, chatID)
	}
}


func HandleBookingsCommand(bot *tgbotapi.BotAPI, chatID int64) {
	bot.Send(tgbotapi.NewMessage(chatID, "Введите имя клиента для получения списка заказов:"))
	UserStates[chatID] = &BookingState{
		Step: 5,
	}
}


func HandleBookingsState(bot *tgbotapi.BotAPI, chatID int64, userMessage string, state *BookingState, bookingService string) {
	clientName := userMessage
	orders, err := GetOrdersForClient(bookingService, clientName)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "Ошибка при получении списка заказов."))
		delete(UserStates, chatID)
		return
	}

	if len(orders) == 0 {
		bot.Send(tgbotapi.NewMessage(chatID, "Список заказов пуст."))
	} else {
		message := "Список ваших заказов:\n"
		for _, order := range orders {
			message += fmt.Sprintf(
				"ID заказа: %d\nОтель: %s\nДата заезда: %s\nДата выезда: %s\n",
				order.ID,
				order.Hotel,
				order.CheckIn.Format("2006-01-02"),
				order.CheckOut.Format("2006-01-02"),
			)
		}
		bot.Send(tgbotapi.NewMessage(chatID, message))
	}

	delete(UserStates, chatID)
}
