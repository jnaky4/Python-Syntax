package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	FIBER_PORT = "fiber.port"
	FIBER_ENV = "fiber.env"
	CORS_ALLOW_HEADERS = "cors.allowHeaders"
	CORS_ALLOW_ORIGINS = "cors.allowOrigins"
)

type FiberConfig struct {
	Config
}

func (c FiberConfig) LoadDBConfig(){
	c.LoadConfig()
	c.CheckRequiredConfig()
}

func (c FiberConfig) CheckRequiredConfig() {
	required := []string{
		FIBER_PORT,
		FIBER_ENV,
		CORS_ALLOW_HEADERS,
		CORS_ALLOW_ORIGINS,
	}
	for _, v := range required {
		if viper.GetString(v) == "" {
			log.Fatalf("Required config '%v' is not defined.\n", v)
		}
	}
}

func NewFiberConfig(configName string) FiberConfig{
	return FiberConfig{ Config: Config{FileName: configName}}
}


