package msger

import (
	"context"
	"log"

	"github.com/MatveyShel/Go_final/pkg/pb"
)

// type RpcMessager interface {
// 	AskPrice(hotelName string) int
// }

type RpcMessager struct {
	rpcClient pb.MessengerServerClient
}

func NewRpcMessager(rpcClient pb.MessengerServerClient) *RpcMessager {
	return&RpcMessager{
		rpcClient: rpcClient,
	}
}

func (m *RpcMessager) AskPrice(ctx context.Context, hotel string) (int, error) {
	r, err := m.rpcClient.AskPrice(ctx, &pb.AskRequest{Hotel: hotel})
	if err != nil {
		log.Fatalf("could not greet: %v", err)

		return 0, err
	}
	return int(r.Price), nil
}