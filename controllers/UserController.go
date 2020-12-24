package controllers

import (
	"encoding/json"
	"go-orderfood/models"
	"go-orderfood/services"
	"go-orderfood/responses"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// Define a struct to return when user is authorized
//@Title Login
//@Description Login
//@Summary "Login"
// @Params username body models.Authorize true
//@Success 200
//@Failure 404
//@router /login [post]
func (this *UserController) Login() {
	defer this.ServeJSON()
	var auth models.Authorize
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &auth)
	if err != nil {
		this.Data["json"] = ""
		return
	}
	cusIns := models.Account{
		Username: auth.Username,
	}
	cus, err := cusIns.Getaccount()
	if err != nil {
		this.Data["json"] = "sai"
		return
	}
	if auth.Password == cus.Password {
		this.Data["json"], _ = services.CreateToken(cus.Typeid, cus.Type)
		return
	}
	this.Data["json"] = ""
	return
}

// Define a struct to return user
//@Title Signup
//@Description Signup
//@Summary "Signup"
// @Params username body true
//@Success 200 string
//@Failure 404 string
//@router /signup [post]
func (this *UserController) Signup(){
	defer this.ServeJSON()
	var acc models.Account

	// this.Data["json"] = ""
	// err := json.Unmarshal(this.Ctx.Input.RequestBody)
	// if err != nil {
	// 	this.Data["json"] = ""
	// 	return
	// }
	
	// if err != nil {
	// 	this.Data["json"] = "sai"
	// 	return
	// }
	// this.Data["json"] = ""
	return
}

// Define a struct to return all Customer
//@Title Get All Customer
//@Description Get All Customer
//@Summary "Get toàn bộ thông tin khách hàng"
// @Params auth header string true "token"
//@Success 200 {object} responses.ResCustomer
//@Failure 404 responses.ResCustomer
//@router /customer [get]
func (this *UserController) GetCustomer() (){
	defer this.ServeJSON()
	var customer models.Customers
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0"{
		this.Data["json"] = responses.ResCustomer{
			Data:  models.Customers{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	customers ,err := customer.GetTotalCustomer()
	if err != nil {
		this.Data["json"] = responses.ResCustomer{
			Data:  models.Customers{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	this.Data["json"] = responses.ResCustomer{
		Data:  customers,
		Error: responses.NewErr(responses.Success),
	}
}

// Define a struct to return all Store
//@Title Get All Store
//@Description Get All Store
//@Summary "Get toàn bộ thông tin Cửa hàng"
// @Params auth header string true "token"
//@Success 200 {object} responses.ResStore
//@Failure 404 responses.ResStore
//@router / [get]
func (this *UserController) GetStore() (){
	defer this.ServeJSON()
	var store models.Store
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0"{
		this.Data["json"] = responses.ResStore{
			Data:  models.Store{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	stores ,err := store.GetTotalStore()
	if err != nil {
		this.Data["json"] = responses.ResStore{
			Data:  models.Store{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	this.Data["json"] = responses.ResStore{
		Data:  stores,
		Error: responses.NewErr(responses.Success),
	}
}

// Define a struct to return all Shipper
//@Title Get All Shipper
//@Description Get All Shipper
//@Summary "Get toàn bộ thông tin người giao hàng"
// @Params auth header string true "token"
//@Success 200 {object} responses.ResShipper
//@Failure 404 responses.ResShipper
//@router / [get]
func (this * UserController) GetShipper(){
	defer this.ServeJSON()
	var shipper models.Shipper
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0"{
		this.Data["json"] = responses.ResShipper{
			Data:  models.Shipper{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	shippers ,err := shipper.GetTotalShipper()
	if err != nil {
		this.Data["json"] = responses.ResShipper{
			Data:  models.Shipper{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	this.Data["json"] = responses.ResShipper{
		Data:  shippers,
		Error: responses.NewErr(responses.Success),
	}
}