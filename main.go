package main

import (
	"go-orderfood/middlewares"
	"go-orderfood/models"
	_ "go-orderfood/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.BConfig.WebConfig.AutoRender = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/v1/swagger"] = "swagger"

	models.InitConnectDataBase()

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Connection", "Authorization", "Sec-WebSocket-Extensions", "Sec-WebSocket-Key",
			"Sec-WebSocket-Version", "Access-Control-Allow-Origin", "content-type", "Content-Type", "sessionkey", "token", "Upgrade"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type", "Sec-WebSocket-Accept", "Connection", "Upgrade"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("*", beego.BeforeRouter, middlewares.Log)

	beego.Run()
}
