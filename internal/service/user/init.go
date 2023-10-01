package user

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	pb "studysystem_micro/idl/user"
	gen "studysystem_micro/internal/service/user/config"
	"studysystem_micro/pkg/init/config"
	"studysystem_micro/pkg/init/log"
	"studysystem_micro/pkg/init/serviceCenter"
	"studysystem_micro/pkg/init/sql"
	"studysystem_micro/pkg/init/tracing"
	"studysystem_micro/pkg/utils"
	"syscall"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func init() {
	gen.C = new(gen.Config)
	config.Init_config("user", "studysystem", gen.C)
	sql.InitMysql(gen.C.Mysql)
	sql.InitRedis(gen.C.Redis)
	log.InitLogger(gen.C.Server.Name)
	log.GetLogger().Debug("Debug log enabled")

}
func Run() {
	logs := log.GetSugarLogger()

	// 获取本机IP
	ipObj, _ := utils.GetOutboundIP()
	ip := ipObj.String()
	//链路追踪
	tracer, closer := tracing.InitTracer(gen.C.Server.Name, gen.C.Tracing.Host, gen.C.Tracing.Port)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	// 创建 Tcp 连接
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", gen.C.Server.Port))
	if err != nil {
		logs.Errorln("服务连接失败", err)
		return
	}
	//开启grpc服务
	grpcServer := grpc.NewServer()
	//健康检查
	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthcheck)
	pb.RegisterUserServer(grpcServer, new(UserServer))

	// 注册服务
	consul, err := serviceCenter.NewConsul(gen.C.Consul.Host + ":" + gen.C.Consul.Port)
	if err != nil {
		logs.Errorln("服务注册失败", err)
		return
	}
	consul.RegisterService(gen.C.Server.Name, ip, gen.C.Server.Port)
	// 启动服务
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			logs.Errorln("服务启动失败", err)
			return
		}
	}()
	// 等待信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	// 退出时注销服务
	consul.Deregister(gen.C.Server.Name, ip, gen.C.Server.Port)
}
