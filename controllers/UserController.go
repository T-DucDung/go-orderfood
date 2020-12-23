package controllers

import (
	"encoding/json"
	"go-orderfood/models"
	"go-orderfood/services"

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
// @Params username body
//@Success 200
//@Failure 404
//@router /signup [post]
func (this *UserController) Signup(){
	defer this.ServeJSON()
	
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