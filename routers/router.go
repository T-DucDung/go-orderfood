package routers

import (
	"orderfood/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1/orderfood",
		beego.NSNamespace("/statisticstore",
			beego.NSInclude(
				&controllers.StatisticStoreController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
