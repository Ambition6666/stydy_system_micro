package user

import (
	"log"
	"net"
	"os"
	"os/signal"
	pb "studysystem_micro/idl/user"
	"studysystem_micro/pkg/init/config"
	"studysystem_micro/pkg/init/serviceCenter"
	"studysystem_micro/pkg/utils"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func init() {
	C = new(Config)
	config.Init_config("user", "studysystem", C)

}
func Run() {
	// 创建 Tcp 连接
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	grpcServer := grpc.NewServer()
	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthcheck)

	pb.RegisterUserServer(grpcServer, new(UserServer))
	consul, err := serviceCenter.NewConsul(C.Consul.Host + ":" + C.Consul.Port)
	if err != nil {
		log.Printf("NewConsul failed, err:%v\n", err)
	}
	ipObj, _ := utils.GetOutboundIP()
	ip := ipObj.String()
	consul.RegisterService("user", ip, 9091)
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Printf("failed to serve: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	// 退出时注销服务
	consul.Deregister("identify")
}
