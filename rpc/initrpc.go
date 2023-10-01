package rpc

import (
	"fmt"
	"studysystem_micro/pkg/init/config"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var conns []*grpc.ClientConn

type run struct {
	err error
}

// 运行代码，不出现错误时才运行
func (r *run) Run(fun func() (*grpc.ClientConn, error)) {
	if r.err != nil {
		return
	}
	conn, err := fun()
	r.err = err
	conns = append(conns, conn)
}

func (r *run) Error() error {
	return r.err
}

// 初始化所有Client
func InitGRPCClients() error {
	C = new(ServerConfig)
	config.Init_config("gateway", "studysystem", C)
	r := new(run)
	r.Run(InitUserGRPC)
	return r.Error()
}

// 关闭所有Client
func CloseGPRCClients() {
	for _, conn := range conns {
		_ = conn.Close()
	}
}

// 创建GRPC客户端
func NewGRPCClient(addr, name string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("consul://%s/%s?healthy=true", addr, name), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), grpc.WithChainUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		return conn, err
	}
	return conn, nil
}
