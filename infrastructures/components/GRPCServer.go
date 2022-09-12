package components

import (
	"arman-estimation-service/types/aliases"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func InitializeGRPCServer(config aliases.StringMap) {
	lis, err := net.Listen("tcp", getAddress())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)

	serviceFunc, ok := config["grpc_services"]
	if !ok {
		panic("failed to find grpc services!")
	}

	serve, ok := serviceFunc.(func(registrar grpc.ServiceRegistrar))
	if !ok {
		panic("invalid grpc server config!")
	}

	// Register all RPC servers on main gRPC server
	serve(grpcServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}

//getAddress get host and port from env variables.
func getAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("GRPC_HOST"), viper.GetInt("GRPC_PORT"))
}
