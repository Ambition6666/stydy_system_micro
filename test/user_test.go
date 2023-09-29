package test

import (
	"context"
	"fmt"
	pb "studysystem_micro/idl/user"
	"testing"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestUser(t *testing.T) {
	conn, err := grpc.Dial("consul://192.168.1.75:8500/user?healthy=true", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	v, _ := client.Register(context.Background(), &pb.UserRegister{
		Email: "123",
	})
	fmt.Println(v)
}
