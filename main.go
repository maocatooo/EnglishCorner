package main

import (
	"EnglishCorner/src/server"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		server.RunServer()
	}
	switch args[0] {
	case "run":
		server.RunServer()
	case "init":
		server.InitData()
	default:
		fmt.Println("error args has in ['run', 'init']")
	}

}
