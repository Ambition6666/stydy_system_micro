package main

import (
	"studysystem_micro/internal/gateway/route"
	"studysystem_micro/rpc"
)

func main() {
	rpc.InitGRPCClients()
	r := route.InitRouter()
	r.Run(":9090")
}
