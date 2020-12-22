package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-orderfood/controllers:CategoryController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:FoodController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:FoodController"],
        beego.ControllerComments{
            Method: "GetFood",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:FoodController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:FoodController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:FoodController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:FoodController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:FoodController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:FoodController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:MainController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:MenuController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:MenuController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:MenuController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:MenuController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:MenuController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:MenuController"],
        beego.ControllerComments{
            Method: "GetMenu",
            Router: "/detail",
            AllowHTTPMethods: []string{"get"},
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

    beego.GlobalControllerRouter["go-orderfood/controllers:StoreController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:StoreController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:StoreController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:StoreController"],
        beego.ControllerComments{
            Method: "Comment",
            Router: "/comment",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
