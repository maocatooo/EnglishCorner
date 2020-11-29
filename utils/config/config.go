package config

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
)

var (
	cfgApplication *viper.Viper
	cfgDatabase    *viper.Viper
	cfgLogs        *viper.Viper
	cfgRedis       *viper.Viper
)

func SetMode() {
	mode := cfgApplication.GetString("mode")
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
func RegisterConfig(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Read config file fail: %s", err.Error())
		return
	}
	err = viper.ReadConfig(bytes.NewBuffer(content))
	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("config not found settings.application")
	}
	{
		cfgLogs = viper.Sub("settings.logs")
		if cfgLogs == nil {
			panic("config not found settings.logs")
		}
		LogConf = InitLog(cfgLogs)
	}
	{
		cfgDatabase = viper.Sub("settings.database")
		if cfgDatabase == nil {
			panic("config not found settings.database")
		}
		DateBaseConf = InitDatabase(cfgDatabase)
	}
	{

		cfgRedis = viper.Sub("settings.redis")
		if cfgRedis == nil {
			panic("config not found settings.redis")
		}
		RedisConf = InitRedisConfig(cfgRedis)
	}
}
