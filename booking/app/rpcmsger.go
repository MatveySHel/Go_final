package app

import "context"

type RpcMessager interface {
	AskPrice(ctx context.Context, hotelName string) (int, error)
}