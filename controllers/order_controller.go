package controllers

import (
	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

// //@Title Create Order
// //@Description Create Order
// //@Summary "Create Order"
// // Param data body requests.ReqOrder true "Req Order"
// //@Success 200 responses.Response
// //@Failure 404 responses.Response
// //@router / [post]
// func (this *OrderController) CreateOrder() {
// 	defer this.ServeJSON()
// 	c := models.Category{}
// 	ls, err := c.GetCa()
// 	if err != nil {
// 		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
// 		return
// 	}
// 	this.Data["json"] = responses.Response{Data: ls, Error: &responses.Err{Code: 200, Message: "Error"}}
// }
