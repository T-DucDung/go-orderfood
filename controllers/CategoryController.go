package controllers

import (
	"fmt"
	"go-orderfood/models"
	"go-orderfood/responses"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

//@Title Get category
//@Description Get category
//@Success 200 string
//@Failure 400 string
//@router / [get]
func (c *CategoryController) Get() {
	defer c.ServeJSON()
	db := models.GetDatabase()
	data, err := db.Query("SELECT * FROM category")
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	categories := make([]models.Category, 0)
	var category = models.Category{}
	for data.Next() {
		err = data.Scan(&category.CreateDate, &category.UpdateDate, &category.ID, &category.Name, &category.IsActive)
		if err != nil {
			fmt.Println(err)
			c.Data["json"] = responses.FailResponse
			return
		}
		categories = append(categories, category)
	}
	c.Data["json"] = responses.Response{Code: 200, Data: categories}
	return
}
