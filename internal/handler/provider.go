package handler

import (
	"github.com/google/wire"
	"github.com/whj1990/bff-grpc/rpc"
	"github.com/whj1990/mine-grrpc/pbs"
)

var ProviderSet = wire.NewSet(rpc.NewGrpcClient, NewClientHandler)

func NewClientHandler(client pbs.HandleServerClient) ClientHandler {
	return &clientHandler{client}
}
