package controllers

import (
	"go-orderfood/models"
	"go-orderfood/responses"
	"go-orderfood/services"
	"log"

	"github.com/astaxie/beego"
)

type ListOrderController struct {
	beego.Controller
}

//@Title Get List Order
//@Description Get List Order
//@Summary "Get danh sách order "
// @Params auth header string true "token"
//@Success 200 {object} responses.ResListOrder
//@Failure 404 responses.ResListOrder
//@router / [get]
func (this *ListOrderController) GetListOrder() {
	defer this.ServeJSON()
	sid := this.Ctx.Request.Header.Get("Id")
	typeid := this.Ctx.Request.Header.Get("Type")
	log.Println(typeid)

	if typeid == "3" {
		o, err := services.GetLsOrderUser(sid)
		if err != nil {
			this.Data["json"] = responses.ResListOrder{
				Data:  []models.Order{},
				Error: responses.NewErr(responses.UnSuccess),
			}
			return
		}

		this.Data["json"] = responses.ResListOrder{
			Data:  o,
			Error: responses.NewErr(responses.Success),
		}
		return
	}
	if typeid == "2" {
		o, err := services.GetLsOrderShipper()
		log.Println(err)
		if err != nil {
			this.Data["json"] = responses.ResListOrder{
				Data:  []models.Order{},
				Error: responses.NewErr(responses.UnSuccess),
			}
			return
		}

		this.Data["json"] = responses.ResListOrder{
			Data:  o,
			Error: responses.NewErr(responses.Success),
		}
		return
	}
	if typeid == "1" {
		o, err := services.GetLsOrderShop(sid)
		if err != nil {
			this.Data["json"] = responses.ResListOrder{
				Data:  []models.Order{},
				Error: responses.NewErr(responses.UnSuccess),
			}
			return
		}

		this.Data["json"] = responses.ResListOrder{
			Data:  o,
			Error: responses.NewErr(responses.Success),
		}
		return
	}

	this.Data["json"] = responses.ResListOrder{
		Data:  []models.Order{},
		Error: responses.NewErr(responses.UnSuccess),
	}
}

//@Title Get List Order Shipper
//@Description Get List Order Shipper
//@Summary "Get danh sách order mà shipper đã nhận"
// @Params auth header string true "token"
//@Success 200 {object} responses.ResListOrder
//@Failure 404 responses.ResListOrder
//@router /shipper [get]
func (this *ListOrderController) GetListOrderShipper() {
	defer this.ServeJSON()
	sid := this.Ctx.Request.Header.Get("Id")
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "2" {
		this.Data["json"] = responses.ResListOrder{
			Data:  []models.Order{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	o, err := services.GetLsOrderShipper1(sid)
	if err != nil {
		this.Data["json"] = responses.ResListOrder{
			Data:  []models.Order{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	this.Data["json"] = responses.ResListOrder{
		Data:  o,
		Error: responses.NewErr(responses.Success),
	}
	return
}

//@Title Up Status Order
//@Description Up Status Order
//@Summary "Shipper nhận đơn"
// @Params auth header string true "token"
// @Params orderid query string true "order id"
//@Success 200 {object} responses.Response
//@Failure 404 responses.Response
//@router / [post]
func (this *ListOrderController) UpStatusOrder() {
	defer this.ServeJSON()
	sid := this.Ctx.Request.Header.Get("Id")
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "2" {
		this.Data["json"] = responses.Response{
			Data:  "Error",
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	orderid := this.GetString("orderid")
	check, err := services.AcceptOrder(sid, orderid)
	if err != nil {
		this.Data["json"] = responses.Response{
			Data:  err.Error(),
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.Response{
		Data:  check,
		Error: responses.NewErr(responses.Success),
	}
}

//@Title End Status Order
//@Description End Status Order
//@Summary "Shipper hoàn thành đơn"
// @Params auth header string true "token"
// @Params orderid query string true "order id"
//@Success 200 {object} responses.Response
//@Failure 404 responses.Response
//@router /shipper [post]
func (this *ListOrderController) EndStatusOrder() {
	defer this.ServeJSON()
	sid := this.Ctx.Request.Header.Get("Id")
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "2" {
		this.Data["json"] = responses.Response{
			Data:  "Error",
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	orderid := this.GetString("orderid")
	check, err := services.EndOrder(sid, orderid)
	if err != nil {
		this.Data["json"] = responses.Response{
			Data:  err.Error(),
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.Response{
		Data:  check,
		Error: responses.NewErr(responses.Success),
	}
}

//@Title Store Accept Order
//@Description Store Accept Order
//@Summary "Store nhận đơn"
// @Params auth header string true "token"
// @Params orderid query string true "order id"
//@Success 200 {object} responses.Response
//@Failure 404 responses.Response
//@router /store [post]
func (this *ListOrderController) StoreUpStatusOrder() {
	defer this.ServeJSON()
	sid := this.Ctx.Request.Header.Get("Id")
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "1" {
		this.Data["json"] = responses.Response{
			Data:  "Error",
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	orderid := this.GetString("orderid")
	check, err := services.StoreAcceptOrder(sid, orderid)
	if err != nil {
		this.Data["json"] = responses.Response{
			Data:  err.Error(),
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.Response{
		Data:  check,
		Error: responses.NewErr(responses.Success),
	}
}
