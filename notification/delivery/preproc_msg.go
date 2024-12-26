package delivery

import (
 "fmt"
 "strings"
)

func FormatBookingMessage(rawMessage string) (string, error) {

	parts := strings.SplitN(rawMessage, ": {", 2)
	if len(parts) < 2 {
	 return "", fmt.Errorf("неверный формат строки")
	}
   

	dataString := strings.TrimSuffix(parts[1], "}")
	dataParts := strings.Split(dataString, ", ")
	bookingData := map[string]string{}
   
	for _, part := range dataParts {
	 keyValue := strings.SplitN(part, ": ", 2)
	 if len(keyValue) == 2 {
	  bookingData[keyValue[0]] = keyValue[1]
	 }
	}
   
	return fmt.Sprintf(
	 "🏨 Бронирование выполнено успешно, подробная информация:\nКлиент: %s\nОтель: %s\nДата заезда: %s\nДата выезда: %s\nОбщая стоимость: %s рублей",
	 bookingData["Client"],
	 bookingData["Hotel"],
	 bookingData["CheckIn"],
	 bookingData["CheckOut"],
	 bookingData["TotalPrice"],
	), nil
}