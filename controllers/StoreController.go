package controllers

import (
	"fmt"
	"go-orderfood/models"
	"go-orderfood/responses"

	"github.com/astaxie/beego"
)

type StoreController struct {
	beego.Controller
}

//@Title Get
//@Description Get
//@Success 200 string
//@Failure 400 string
//@router / [get]
func (c *StoreController) Get() {
	defer c.ServeJSON()
	db := models.GetDatabase()
	data, err := db.Query("SELECT created_date, updated_date, id, name, rate_avg, rate_one, rate_two, rate_three, rate_four, rate_five FROM store WHERE id = " + c.Ctx.Input.Header("StoreID"))
	if err != nil {
		fmt.Println("store controller", err)
		c.Data["json"] = responses.FailResponse
		return
	}
	stores := make([]models.Store, 0)
	var store = models.Store{}
	for data.Next() {
		err = data.Scan(&store.CreateDate, &store.UpdateDate, &store.ID, &store.Name, &store.RateAvg, &store.RateOne, &store.RateTwo, &store.RateThree, &store.RateFour, &store.RateFive)
		if err != nil {
			fmt.Println("store controller", err)
			c.Data["json"] = responses.FailResponse
			return
		}
		stores = append(stores, store)
	}
	c.Data["json"] = responses.Response{Code: 200, Data: stores}
	return
}

//@Title Comment
//@Description Comment
//@Success 200 string
//@Failure 400 string
//@router /comment [get]
func (c *StoreController) Comment() {
	defer c.ServeJSON()
	db := models.GetDatabase()
	data, err := db.Query("SELECT created_date, updated_date, id, store_id, comment, user_id FROM comment WHERE store_id = " + c.Ctx.Input.Header("StoreID"))
	if err != nil {
		fmt.Println("store controller", err)
		c.Data["json"] = responses.FailResponse
		return
	}
	comments := make([]models.Comment, 0)
	var comment = models.Comment{}
	for data.Next() {
		err = data.Scan(&comment.CreateDate, &comment.UpdateDate, &comment.ID, &comment.Comment, &comment.UserID)
		if err != nil {
			fmt.Println("store controller", err)
			c.Data["json"] = responses.FailResponse
			return
		}
		comments = append(comments, comment)
	}
	c.Data["json"] = responses.Response{Code: 200, Data: comments}
	return
}
