package domain

import (
	"time"
)

type Booking struct {
	ID 			int				`db:"id"`
	Client 		string			`db:"client"`
	Hotel 		string			`db:"hotel"`
	CheckIn 	time.Time		`db:"checkin"`
	CheckOut 	time.Time		`db:"checkout"`
}