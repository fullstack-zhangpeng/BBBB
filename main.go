package main

import (
	_ "Service/service-server/config"
	_ "Service/service-server/db"
	"Service/service-server/router"
)

func main() {
	//dt := dingtalk.NewDingTalk("2062dd60bdc9bae2412e289c83d3e3613ecf4205f30df9c93e294cc5e00a5791")
	//message := dingtalk.NewTextMessage("zhangpeng")
	//message.SetAt([]string{"15810746400"}, false)
	//message := dingtalk.NewLinkMessage(dingtalk.Link{
	//	Title:      "zhangpeng",
	//	Text:       "zhangpeng_测试连接",
	//	PicURL:     "https://www.baidu.com/img/bd_logo1.png",
	//	MessageURL: "http://www.baidu.com",
	//})
	//dt.Send(message)

	//cornManager.InitCorn()
	router := router.InitRouter()
	router.Run()
}
