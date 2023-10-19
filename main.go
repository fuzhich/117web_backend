package main

import (
	"117web/config"
	"117web/route"
)

func main() {
	config.Init()
	router := route.InitRouter()
	addr := ":" + config.Config.GetString("system.port")
	router.Run(addr)
}
