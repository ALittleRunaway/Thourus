package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type NatsConfig struct {
	Addr               string
	InviteChannelName  string
	ConfirmChannelName string
}

func InitNatsConfig() *NatsConfig {

	natsConfig := NatsConfig{
		Addr:               fmt.Sprintf("nats://%s", viper.GetString("nats.addr")),
		InviteChannelName:  viper.GetString("nats.invite_channel_name"),
		ConfirmChannelName: viper.GetString("nats.confirm_channel_name"),
	}
	return &natsConfig
}

func init() {
	viper.SetDefault("nats.addr", "localhost:4222")
	viper.SetDefault("nats.invite_channel_name", "thourus.mailman.invite")
	viper.SetDefault("nats.confirm_channel_name", "thourus.mailman.confirm")

	InitError(viper.BindEnv("nats.addr", EnvPrefix+"_NATS_ADDR"))
	InitError(viper.BindEnv("nats.invite_channel_name", EnvPrefix+"_NATS_INVITE_CHANNEL_NAME"))
	InitError(viper.BindEnv("nats.confirm_channel_name", EnvPrefix+"_NATS_CONFIRM_CHANNEL_NAME"))
}
