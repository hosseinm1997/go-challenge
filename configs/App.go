package configs

import (
	estimation "arman-estimation-service/grpc/estimation/v1"
	segment "arman-estimation-service/grpc/segment/v1"
	"arman-estimation-service/providers"
	"arman-estimation-service/types/aliases"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func App() aliases.StringMap {
	return aliases.StringMap{
		"environment": viper.GetString("APP_ENV"),
		// providers used for registering services into service container (IoC)
		"providers": []func(){
			providers.SystemProvides,
			providers.ServiceProvider,
		},
		// All RPC servers will be registered on main gRPC server by this running this function
		"grpc_services": func(registrar grpc.ServiceRegistrar) {
			segment.RegisterSegmentServer(registrar, segment.SegmentImplServ{})
			estimation.RegisterEstimationServer(registrar, estimation.EstimationImplServ{})
		},
	}
}
