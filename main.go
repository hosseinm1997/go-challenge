package main

import (
	"arman-estimation-service/configs"
	"arman-estimation-service/infrastructures"
)

func main() {
	infrastructures.KernelBuilder().Build(configs.App())
}
