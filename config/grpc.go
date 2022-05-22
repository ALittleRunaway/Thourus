package config

import (
	"github.com/spf13/viper"
)

type GrpcConfig struct {
	Network string
	Addr    string
}

func InitGrpcConfig() *GrpcConfig {

	grpcConfig := GrpcConfig{
		Network: viper.GetString("grpc.network"),
		Addr:    viper.GetString("grpc.addr"),
	}
	return &grpcConfig
}

func init() {
	viper.SetDefault("grpc.network", "tcp")

	InitError(viper.BindEnv("grpc.network", EnvPrefix+"_GRPC_NETWORK_TYPE"))
	InitError(viper.BindEnv("grpc.addr", EnvPrefix+"_GRPC_ADDR"))
}
