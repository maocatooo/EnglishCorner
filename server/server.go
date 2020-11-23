package server

import (
	"EnglishCorner/db"
	"EnglishCorner/helps"
	"EnglishCorner/router"
	"EnglishCorner/utils/config"
	"EnglishCorner/utils/log"
	"fmt"
	"os"
)

func initBase() {

	basePath, _ := os.Getwd()
	config.RegisterConfig(basePath + "/conf/settings.yaml")
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
	helps.InitData(basePath + "/conf/dict.txt")
}

func Import(file string) {
	initBase()
	//helps.DeleteTableData()
	//helps.CreateTable()
	basePath, _ := os.Getwd()
	fmt.Println(file)
	helps.ImportFile(basePath+"/conf/", file)
}
