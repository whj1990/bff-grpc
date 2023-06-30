package rpc

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/whj1990/go-core/config"
	"github.com/whj1990/go-core/handler"
	"github.com/whj1990/mine-grrpc/pbs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"strings"
)

func NewGrpcClient() pbs.HandleServerClient {
	cert, err := credentials.NewClientTLSFromFile("./cert/server.crt", "localhost")
	if err != nil {
		zap.L().Error(err.Error())
	}
	instance := config.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: config.GetNacosConfigData().ClientServer.GrpcServerName,
		GroupName:   config.GetNacosConfigData().ClientServer.GrpcGroupName,                       // 默认值DEFAULT_GROUP
		Clusters:    strings.Split(config.GetNacosConfigData().ClientServer.GrpcClusterName, ","), // 默认值DEFAULT
	})

	//dialAddress := config.GetNacosConfigData().ClientServer.DialAddress
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", instance.Ip, instance.Port),
		[]grpc.DialOption{
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
			grpc.WithTransportCredentials(cert),
		}...)
	if err != nil {
		handler.HandleError(err)
	}
	client := pbs.NewHandleServerClient(conn)
	return client
}
