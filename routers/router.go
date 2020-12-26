package routers

import (
	"go-orderfood/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/statisticstore",
			beego.NSInclude(
				&controllers.StatisticStoreController{},
			),
		),
		beego.NSNamespace("/statistic",
			beego.NSInclude(
				&controllers.StatisticController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/food",
			beego.NSInclude(
				&controllers.FoodController{},
			),
		),
		beego.NSNamespace("/listorder",
			beego.NSInclude(
				&controllers.ListOrderController{},
			),
		),
		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),
		beego.NSNamespace("/store",
			beego.NSInclude(
				&controllers.StoreController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
