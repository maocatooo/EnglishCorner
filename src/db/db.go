package db

import (
	"EnglishCorner/src/utils/log"
	//"gorm.io/driver/sqlite"
	"EnglishCorner/src/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	var err error
	dns := config.DateBaseConf.DNS()
	log.Debug(dns)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Error("数据库链接失败")
		panic("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return db
}
