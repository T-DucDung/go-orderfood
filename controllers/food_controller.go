package controllers

import (
	"go-orderfood/models"
	"go-orderfood/responses"

	"github.com/astaxie/beego"
)

type FoodController struct {
	beego.Controller
}

//@Title Get Category
//@Description Get Category
//@Summary "Get Category"
//@Success 200 responses.Response
//@Failure 404 responses.Response
//@router /cate [get]
func (this *FoodController) GetCategory() {
	defer this.ServeJSON()
	c := models.Category{}
	ls, err := c.GetCa()
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	this.Data["json"] = responses.Response{Data: ls, Error: &responses.Err{Code: 200, Message: "Error"}}
}

//@Title Get List Food Rand
//@Description Get List Food Rand
//@Summary "Get List Food Rand"
//@Success 200 responses.Response
//@Failure 404 responses.Response
//@router /listrand [get]
func (this *FoodController) GetListFoodRand() {
	defer this.ServeJSON()
	c := models.Food{}
	ls, err := c.GetRandFood()
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	this.Data["json"] = responses.Response{Data: ls, Error: &responses.Err{Code: 200, Message: "Error"}}
}

//@Title Get List Food Rand By Cate
//@Description Get List Food Rand By Cate
//@Summary "Get List Food Rand By Cate"
//@Success 200 responses.Response
//@Failure 404 responses.Response
//@router /listrandbycate [get]
func (this *FoodController) GetListFoodRandByCate() {
	defer this.ServeJSON()
	c := models.Food{}
	ls, err := c.GetRandfoodCate()
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	this.Data["json"] = responses.Response{Data: ls, Error: &responses.Err{Code: 200, Message: "Error"}}
}
