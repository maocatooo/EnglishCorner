package db

import (
	"EnglishCorner/utils/config"
	"EnglishCorner/utils/log"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool //创建redis连接池

func InitRedis() {

	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			log.Error(config.RedisConf.Url)
			c, e := redis.Dial("tcp", config.RedisConf.Url)
			if e != nil {
				log.Error("redis链接失败")
				panic("failed to connect redis")
			}
			return c, e
		},
	}
}

func GetRedisConn() redis.Conn {
	return pool.Get()
}
