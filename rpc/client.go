package rpc

import (
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/whj1990/go-core/config"
	"github.com/whj1990/go-core/handler"
	"github.com/whj1990/mine-grrpc/pbs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewGrpcClient() pbs.HandleServerClient {
	cert, err := credentials.NewClientTLSFromFile("./cert/server.crt", "localhost")
	if err != nil {
		zap.L().Error(err.Error())
	}
	dialAddress := config.GetNacosConfigData().ClientServer.DialAddress
	conn, err := grpc.Dial(dialAddress,
		[]grpc.DialOption{
			//grpc.WithBlock(),
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
			grpc.WithTransportCredentials(cert),
		}...)
	if err != nil {
		handler.HandleError(err)
	}
	client := pbs.NewHandleServerClient(conn)
	return client
}
