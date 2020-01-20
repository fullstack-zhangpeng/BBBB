package main

import (
	"MyService/router"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	router := router.InitRouter()
	router.Run()
}
