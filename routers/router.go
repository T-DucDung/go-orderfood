package routers

import (
	"orderfood/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Get("/", func(ctx *context.Context) {
		_ = ctx.Output.Body([]byte("This is a Beego + JWT API - Creator: Mehran Abghari (mehran.ab80@gmail.com)"))
	})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/statisticstore",
			beego.NSInclude(
				&controllers.StatisticStoreController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
