package controllers

import (
	"encoding/json"
	"go-orderfood/models"
	"go-orderfood/requests"
	"go-orderfood/responses"
	"log"
	"go-orderfood/services"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//@Title Login
//@Description Login
//@Summary "Login"
// @Params username body models.Authorize true "username"
//@Success 200
//@Failure 404
//@router / [post]
func (this *LoginController) Login() {
	defer this.ServeJSON()
	var auth models.Authorize
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &auth)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 400, Message: "Error"}}
		return
	}
	cusIns := models.Account{
		Username: auth.Username,
	}
	cus, err := cusIns.Getaccount()
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 404, Message: "UserName Not Found"}}
		return
	}
	if auth.Password == cus.Password {
		token, _ := services.CreateToken(cus.Typeid, cus.Type)
		this.Data["json"] = responses.Response{Data: token, Error: &responses.Err{Code: 200, Message: "Success"}}
		return
	}
	this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 400, Message: "Error"}}
	return
}


//@Title Signup
//@Description Signup
//@Summary "Signup"
// @Params username body requests.CreateCustomer true "username"
//@Success 200
//@Failure 404
//@router /signup [post]
func (this *LoginController) Signup() {
	defer this.ServeJSON()
	var cus models.Customers

	req := requests.CreateCustomer{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Data["json"] = responses.ResCustomer{
			Data:  []models.Customers{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	cusIns := models.Account{
		Username: req.Username,
	}
	_, err = cusIns.Getaccount()
	if err == nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 400, Message: "UserName already exists"}}
		return
	}
	
	_, err = cus.CreateCustomer(req)
	if err != nil {
		this.Data["json"] = responses.ResCustomer{
			Data:  []models.Customers{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	log.Println(err)
	this.Data["json"] = responses.ResCustomer{
		Data:  []models.Customers{},
		Error: responses.NewErr(responses.Success),
	}
}


