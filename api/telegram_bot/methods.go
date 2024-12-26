package telegram_bot

import (
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/MatveyShel/Go_final/hotels/domain"
	dmb "github.com/MatveyShel/Go_final/booking/domain"
)



func GetOrdersForClient(booking_service, clientName string) ([]dmb.Booking, error) {

    requestBody, err := json.Marshal(map[string]string{
        "client": clientName,
    })
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("GET", "http://"+booking_service+"/booking", bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")

    httpClient := &http.Client{}
    resp, err := httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }
    var bookings []dmb.Booking
    decoder := json.NewDecoder(resp.Body)
    if err := decoder.Decode(&bookings); err != nil {
        return nil, err
    }

    return bookings, nil
}

func GetHotels(hotel_service string) (string, error) {
	resp, err := http.Get("http://"+hotel_service+"/hotels")
	if err != nil {
	 return "", err
	}
	defer resp.Body.Close()
   
	if resp.StatusCode != http.StatusOK {
	 return "", fmt.Errorf("не удалось получить данные об отелях, статус ответа: %d", resp.StatusCode)
	}
   
	var hotels []domain.Hotel
   
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&hotels); err != nil {
	 return "", err
	}
   
	var result string
	for _, hotel := range hotels {
	 result += fmt.Sprintf("ID: %d\nНазвание: %s\nЦена: %d\nГород: %s\n\n", hotel.ID, hotel.Name, hotel.Price, hotel.City)
	}
	return result, nil
   }


type Booking struct {
	Client   string `json:"client"`
	Hotel    string `json:"hotel"`
	CheckIn  string `json:"checkin"`
	CheckOut string `json:"checkout"`
}

func BookHotel(booking_service string, booking Booking) (string, error) {


	jsonData, err := json.Marshal(booking)
	if err != nil {
		return "", fmt.Errorf("ошибка при сериализации бронирования: %v", err)
	}

	resp, err := http.Post("http://"+booking_service+"/booking", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ошибка бронирования, статус ответа: %d", resp.StatusCode)
	}

	var responseMessage struct {
		Message string `json:"message"`
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&responseMessage); err != nil {
		return "", fmt.Errorf("ошибка при декодировании ответа: %v", err)
	}

	return responseMessage.Message, nil
}
