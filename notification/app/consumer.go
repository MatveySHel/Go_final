package app

import (
	"context"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type Consumer interface {
	StartConsuming(ctx context.Context)
	Close() error
	GetchatID() int64
	Getbot() *tgbotapi.BotAPI
}
