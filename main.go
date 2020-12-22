package main

import (
	"orderfood/models"
	_ "orderfood/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.AutoRender = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/v1/orderfood/swagger"] = "swagger"

	models.InitConnectDataBase()
	beego.Run()
}
