package config

import (
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Url string //日志存路径
}

func InitRedisConfig(cfg *viper.Viper) *RedisConfig {
	return &RedisConfig{
		Url: cfg.GetString("url"),
	}
}

var RedisConf = new(RedisConfig)
