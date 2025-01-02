package main

import configs "github.com/CrisPassos/goexpert/7_APIs/config"

func main() {
	// Start the server
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
