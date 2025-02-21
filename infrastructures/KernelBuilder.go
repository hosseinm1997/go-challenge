package infrastructures

import (
	"arman-estimation-service/types/aliases"
	"sync"
)

//KernelBuilder builds kernel components one by one in a correct order. It knows when to build each component.
// It uses "Buildr Pattern"

type IKernelBuilder interface {
	Build(config aliases.StringMap) IKernel
}

type kernelBuilder struct{}

var (
	k          *kernel
	kernelOnce sync.Once
)

func (kb *kernelBuilder) Build(config aliases.StringMap) IKernel {
	if k == nil {
		kernelOnce.Do(func() {
			k = &kernel{}
			k.loadEnvVars()
			k.initServiceContainer(config)
			k.initRedis(config)
			k.initGRPCServer(config)
		})
	}

	return k
}

func KernelBuilder() IKernelBuilder {
	return &kernelBuilder{}
}
