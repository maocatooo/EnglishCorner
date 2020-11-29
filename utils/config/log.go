package config

import (
	"github.com/spf13/viper"
	"strings"
)

type LogConfig struct {
	LogDir string //日志存路径
}

func InitLog(cfg *viper.Viper) *LogConfig {
	ld := cfg.GetString("logdir")
	if !strings.HasSuffix(ld, "/") {
		ld = ld + "/"
	}
	return &LogConfig{
		LogDir: ld,
	}
}

var LogConf = new(LogConfig)
