package main

import "github.com/CrisPassos/goexperts/08_apis/configs"

func main() {
	config, _ := configs.LoadConfig("./")
	println(config.DBDrive)
}
