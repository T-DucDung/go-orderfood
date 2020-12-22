package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["orderfood/controllers:StatisticStoreController"] = append(beego.GlobalControllerRouter["orderfood/controllers:StatisticStoreController"],
        beego.ControllerComments{
            Method: "GetStatisticStore",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
