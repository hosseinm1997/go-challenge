package infrastructures

//Kernel is the starting point of this microservice. It runs each component starting point.

import (
	"arman-estimation-service/infrastructures/components"
	"arman-estimation-service/types/aliases"
)

type IKernel interface {
	loadEnvVars()
	initServiceContainer(config aliases.StringMap)
	initGRPCServer(config aliases.StringMap)
	initRedis(config aliases.StringMap)
}

type kernel struct{}

func (k *kernel) loadEnvVars() {
	if err := components.InitEnvLoader().LoadFromFile(".env"); err != nil {
		println(err)
	}
}

func (k *kernel) initServiceContainer(config aliases.StringMap) {
	components.InitializeServiceContainer(config)
}

func (k *kernel) initGRPCServer(config aliases.StringMap) {
	components.InitializeGRPCServer(config)
}

func (k *kernel) initRedis(config aliases.StringMap) {
	components.InitializeRedisServer(config)
}
