package rpc

import (
	pb "studysystem_micro/idl/user"
	
	"google.golang.org/grpc"
	
)

var userGrpcCli pb.UserClient

func GetUserCli()*pb.UserClient{
	return &userGrpcCli
}

func InitUserGRPC() (*grpc.ClientConn, error) {
	conn, err := NewGRPCClient(C.Consul.Host+":"+C.Consul.port, C.User.Name)
	if err != nil {
		return conn, err
	}
	userGrpcCli=pb.NewUserClient(conn)

	return conn, nil
}