package server

import (
	"EnglishCorner/src/db"
	"EnglishCorner/src/helps"
	"EnglishCorner/src/router"
	"EnglishCorner/src/utils/config"
	"EnglishCorner/src/utils/log"
	"os"
)

func initBase() {

	basePath, _ := os.Getwd()
	config.RegisterConfig(basePath + "/src/conf/settings.yaml")
	config.SetMode()
	log.InitLogger()
	db.InitDB()

}
func RunServer() {
	initBase()
	helps.CreateTable()
	r := router.InitRouter()
	r.Run(":8011")
}

func InitData() {
	initBase()
	helps.DeleteTableData()
	helps.CreateTable()
	basePath, _ := os.Getwd()
	helps.InitData(basePath + "/src/conf/dict.text")
}
