package main

import (
	"github.com/whj1990/bff-grpc/internal/router"
	"github.com/whj1990/go-core/launch"
)

// @title			Swagger bff-grpc
// @version		1.0
// @license.name Apache 2.0
func main() {
	logger, closer := launch.InitPremise(false)
	defer logger.Sync()
	defer closer.Close()
	launch.InitHttpServer(initServer().Routes...)
}

func newRouterQuote(routeHandler *router.RouteHandler) *launch.RouterQuote {
	return &launch.RouterQuote{Routes: []launch.HttpRouter{
		routeHandler,
	}}
}
