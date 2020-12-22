package controllers

import (
	"go-orderfood/models"
	"go-orderfood/responses"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//@Title Get
//@Description Get
//@Success 200 string
//@Failure 400 string
//@router / [get]
func (c *MainController) Get() {
	defer c.ServeJSON()
	db := models.GetDatabase()
	data, err := db.Query("SELECT * FROM store")
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	stores := make([]models.Store, 0)
	var store = models.Store{}
	for data.Next() {
		err = data.Scan(&store.CreateDate, &store.UpdateDate, &store.ID, &store.Name, &store.RateAvg, &store.RateOne, &store.RateTwo, &store.RateThree, &store.RateFour, &store.RateFive)
		if err != nil {
			c.Data["json"] = responses.FailResponse
			return
		}
		stores = append(stores, store)
	}
	c.Data["json"] = responses.Response{Code: 200, Data: stores}
	return
}
