package main

import (
	_ "Service/service-server/config"
	_ "Service/service-server/db"
	"Service/service-server/router"
	"Service/service-server/timing"
)

func main() {
	timing.Init()
	timing.Start()
	router := router.InitRouter()
	router.Run()
}
