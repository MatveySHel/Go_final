package telegram_bot

var UserStates = make(map[int64]*BookingState)

type BookingState struct {
	Step    int
	Booking Booking
}