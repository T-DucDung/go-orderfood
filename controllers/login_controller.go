package controllers

import (
	"encoding/json"
	"go-orderfood/models"
	"go-orderfood/responses"
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
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	cusIns := models.Account{
		Username: auth.Username,
	}
	cus, err := cusIns.Getaccount()
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	if auth.Password == cus.Password {
		token, _ := services.CreateToken(cus.Typeid, cus.Type)
		this.Data["json"] = responses.Response{Data: token, Error: &responses.Err{Code: 200, Message: "Success"}}
		return
	}
	this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
	return
}
