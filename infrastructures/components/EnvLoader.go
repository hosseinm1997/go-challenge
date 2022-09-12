package components

// This component used to load all env variables from .env file
// In the future it will use real environment variables set by server

import (
	"github.com/spf13/viper"
	"sync"
)

type IEnvLoader interface {
	LoadFromFile(filename string) error
}

type envLoader struct{}

func (e *envLoader) LoadFromFile(filename string) error {
	viper.SetConfigFile(filename)

	return viper.ReadInConfig()
}

var (
	el     *envLoader
	elOnce sync.Once
)

func InitEnvLoader() *envLoader {
	if el == nil {
		elOnce.Do(func() {
			el = &envLoader{}
		})
	}

	return el
}
