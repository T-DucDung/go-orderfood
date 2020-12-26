package controllers

import (
	"encoding/json"
	"go-orderfood/models"
	"go-orderfood/requests"
	"go-orderfood/responses"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

//@Title Create Order
//@Description Create Order
//@Summary "Create Order"
// @Params auth header string true "token"
// @Param data body requests.ReqOrder true "Req Order"
//@Success 200 responses.Response
//@Failure 404 responses.Response
//@router / [post]
func (this *OrderController) CreateOrder() {
	defer this.ServeJSON()
	id, err := strconv.Atoi(this.Ctx.Request.Header.Get("id"))
	if err != nil {
		log.Println(err)
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "3" {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	log.Println(id)
	req := requests.ReqOrder{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		log.Println(err)
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	log.Println(req)
	o := models.Order{}
	err = o.CreateOrder(req, strconv.Itoa(id))
	if err != nil {
		log.Println(err)
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 200, Message: "Success"}}
}
