package server

import (
	"EnglishCorner/db"
	"EnglishCorner/helps"
	"EnglishCorner/router"
	"EnglishCorner/utils/config"
	"EnglishCorner/utils/log"
	"os"
)

func initBase() {

	basePath, _ := os.Getwd()
	config.RegisterConfig(basePath + "/conf/settings.yaml")
	config.SetMode()
	log.InitLogger()
	db.InitDB()
	db.InitRedis()

}
func RunServer() {
	initBase()
	//helps.CreateTable()
	r := router.InitRouter()
	_ = r.Run(":8011")
}

func InitData() {
	initBase()
	helps.DeleteTableData()
	helps.CreateTable()
	basePath, _ := os.Getwd()
	helps.InitData(basePath + "/conf/dict.txt")
}

func Import(file string) {
	initBase()
	//helps.DeleteTableData()
	//helps.CreateTable()
	basePath, _ := os.Getwd()
	helps.ImportFile(basePath+"/conf/", file)
}
