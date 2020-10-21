package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DataBaseConfig struct {
	Username string //账号
	Password string //密码
	Host     string //数据库地址，可以是Ip或者域名
	Port     int    //数据库端口
	DBname   string //数据库名
	Timeout  string //连接超时，10秒
}

func (d *DataBaseConfig) DNS() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ss",
		d.Username, d.Password, d.Host, d.Port, d.DBname, d.Timeout)
	return dsn
}
func InitDatabase(cfg *viper.Viper) *DataBaseConfig {
	return &DataBaseConfig{
		Port:     cfg.GetInt("port"),
		Host:     cfg.GetString("host"),
		DBname:   cfg.GetString("dbname"),
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
		Timeout:  "10",
	}
}

var DateBaseConf = new(DataBaseConfig)
