package router

import (
	"github.com/google/wire"
	"github.com/whj1990/bff-grpc/internal/handler"
)

var ProviderSet = wire.NewSet(NewRouteHandler)

func NewRouteHandler(clientHandler handler.ClientHandler) *RouteHandler {
	return &RouteHandler{clientHandler}
}
