package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-orderfood/controllers:ListOrderController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:ListOrderController"],
        beego.ControllerComments{
            Method: "GetListOrder",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:ListOrderController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:ListOrderController"],
        beego.ControllerComments{
            Method: "UpStatusOrder",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:ListOrderController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:ListOrderController"],
        beego.ControllerComments{
            Method: "GetListOrderShipper",
            Router: "/shipper",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:ListOrderController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:ListOrderController"],
        beego.ControllerComments{
            Method: "EndStatusOrder",
            Router: "/shipper",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:ListOrderController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:ListOrderController"],
        beego.ControllerComments{
            Method: "StoreUpStatusOrder",
            Router: "/store",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:StatisticController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:StatisticController"],
        beego.ControllerComments{
            Method: "GetStatistic",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:StatisticStoreController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:StatisticStoreController"],
        beego.ControllerComments{
            Method: "GetStatisticStore",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:UserController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:UserController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:UserController"],
        beego.ControllerComments{
            Method: "Signup",
            Router: "/signup",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
