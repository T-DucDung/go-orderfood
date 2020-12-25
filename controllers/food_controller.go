package controllers

import (
	"fmt"
	"go-orderfood/models"
	"go-orderfood/responses"
	"log"

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

//@Title Get List Food
//@Description Get List Food
//@Summary "Get List Food
// @Params name query string false "name"
// @Params cate query string false "type"
//@Success 200 responses.Response
//@Failure 404 responses.Response
//@router /list [get]
func (this *FoodController) GetListFood() {
	defer this.ServeJSON()
	name := this.GetString("name")
	cate := this.GetString("cate")
	if len(name) == 0 {
		if len(cate) == 0 {
			c := models.Food{}
			ls, err := c.GetFood()
			if err != nil {
				this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
				return
			}
			this.Data["json"] = responses.Response{Data: ls, Error: &responses.Err{Code: 200, Message: "Error"}}
			return
		} else {
			c := models.Food{}
			ls, err := c.GetFoodByCate(cate)
			if err != nil {
				fmt.Println(err)
				this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
				return
			}
			this.Data["json"] = responses.Response{Data: ls, Error: &responses.Err{Code: 200, Message: "Error"}}
			return
		}
		c := models.Food{}
		ls, err := c.GetFoodByCateAndName(cate, name)
		if err != nil {
			this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
			return
		}
		this.Data["json"] = responses.Response{Data: ls, Error: &responses.Err{Code: 200, Message: "Error"}}
		return
	}
	c := models.Food{}
	ls, err := c.GetFoodByName(name)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	this.Data["json"] = responses.Response{Data: ls, Error: &responses.Err{Code: 200, Message: "Error"}}
	return

}

//@Title Get List Food In Store
//@Description Get List Food In Store
//@Summary "Get List Food In Store"
// @Params sid query string true "sid"
//@Success 200 responses.Response
//@Failure 404 responses.Response
//@router /liststore [get]
func (this *FoodController) GetFoodStore() {
	defer this.ServeJSON()
	sid := this.GetString("sid")
	log.Println(sid)
	c := models.Food{}
	ls, err := c.GetListInStore(sid)
	if err != nil {
		log.Println(err)
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	this.Data["json"] = responses.Response{Data: ls, Error: &responses.Err{Code: 200, Message: "Error"}}
}

// //@Title Get List Store
// //@Description Get List Store
// //@Summary "Get List Store"
// //@Success 200 responses.Response
// //@Failure 404 responses.Response
// //@router /store [get]
// func (this *FoodController) GetListStore() {
// 	defer this.ServeJSON()
// 	c := models.Food{}
// 	ls, err := c.GetRandFood()
// 	if err != nil {
// 		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
// 		return
// 	}
// 	this.Data["json"] = responses.Response{Data: ls, Error: &responses.Err{Code: 200, Message: "Error"}}
// }
