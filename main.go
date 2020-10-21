package main

import (
	"EnglishCorner/src/server"
	"flag"
	"fmt"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "", "文件名称,文件需要放在conf下且文件名是词库名称")
	a := flag.String("fa", "", "文件名称,文件需要放在conf下且文件名是词库名称")

	flag.Parse()
	fmt.Println(*a)
	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		server.RunServer()
	}
	switch args[0] {
	case "run":
		server.RunServer()
	case "init":
		server.InitData()
	case "import":
		server.Import(file)
	default:
		fmt.Println("error args has in ['run', 'init']")
	}

}
