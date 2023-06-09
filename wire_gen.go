// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/whj1990/bff-grpc/internal/handler"
	"github.com/whj1990/bff-grpc/internal/router"
	"github.com/whj1990/bff-grpc/rpc"
	"github.com/whj1990/go-core/launch"
)

// Injectors from wire.go:

func initServer() *launch.RouterQuote {
	handleServerClient := rpc.NewGrpcClient()
	clientHandler := handler.NewClientHandler(handleServerClient)
	routeHandler := router.NewRouteHandler(clientHandler)
	routerQuote := newRouterQuote(routeHandler)
	return routerQuote
}
