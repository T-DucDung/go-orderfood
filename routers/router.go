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

		beego.NSNamespace("/food",
			beego.NSInclude(
				&controllers.FoodController{},
			),
		),
		beego.NSNamespace("/menu",
			beego.NSInclude(
				&controllers.MenuController{},
			),
		),
		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
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