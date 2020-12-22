package controllers

import (
	"encoding/json"
	"fmt"
	"go-orderfood/models"
	"go-orderfood/responses"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type FoodController struct {
	beego.Controller
}

//@Title Get food
//@Description Get food
//@Success 200 string
//@Failure 400 string
//@router / [get]
func (c *FoodController) GetFood() {
	defer c.ServeJSON()
	db := models.GetDatabase()
	data, err := db.Query("SELECT created_date,updated_date,id,name,init_price,sale_price,title,content,store_id,menu_id,category_id,status,image FROM food WHERE store_id = " + c.Ctx.Input.Header("StoreID"))
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	foods := make([]models.Food, 0)
	var food = models.Food{}
	for data.Next() {
		err = data.Scan(&food.CreateDate, &food.UpdateDate, &food.ID, &food.Name, &food.InitPrice, &food.SalePrice, &food.Title, &food.Content, &food.StoreID, &food.MenuID, &food.CategoryID, &food.Status, &food.Image)
		if err != nil {
			fmt.Println(err)
			c.Data["json"] = responses.FailResponse
			return
		}
		foods = append(foods, food)
	}
	c.Data["json"] = responses.Response{Code: 200, Data: foods}
	return
}

//@Title Post food
//@Description Post
//@Params body body models.Food true "food"
//@Success 200 string
//@Failure 400 string
//@router / [post]
func (c *FoodController) Post() {
	defer c.ServeJSON()
	var food models.Food
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &food)
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	food.CreateDate = time.Now().Unix()
	food.UpdateDate = time.Now().Unix()
	food.StoreID, err = strconv.Atoi(c.Ctx.Input.Header("StoreID"))
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	db := models.GetDatabase()

	data, err := db.Prepare("INSERT INTO food (created_date, updated_date, name, init_price, sale_price, title, content, store_id, menu_id, category_id, status, image) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		panic(err.Error())
	}
	result, err := data.Exec(food.CreateDate, food.UpdateDate, food.Name, food.InitPrice, food.SalePrice, food.Title, food.Content, food.StoreID, food.MenuID, food.CategoryID, food.Status, food.Image)
	fmt.Println(result)
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	c.Data["json"] = responses.Response{Code: 200, Data: food}
	return
}

//@Title Put food
//@Description Put
//@Params body body models.Food true "food"
//@Success 200 string
//@Failure 400 string
//@router / [put]
func (c *FoodController) Put() {
	defer c.ServeJSON()
	var food models.Food
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &food)
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	food.UpdateDate = time.Now().Unix()
	food.StoreID, err = strconv.Atoi(c.Ctx.Input.Header("StoreID"))
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	db := models.GetDatabase()

	data, err := db.Prepare("UPDATE food SET created_date=?, updated_date=?, name=?, init_price=?, sale_price=?, title=?, content=?, store_id=?, menu_id=?, category_id=?, status=?, image=? WHERE id=?;")
	if err != nil {
		panic(err.Error())
	}
	result, err := data.Exec(food.CreateDate, food.UpdateDate, food.Name, food.InitPrice, food.SalePrice, food.Title, food.Content, food.StoreID, food.MenuID, food.CategoryID, food.Status, food.Image, food.ID)
	fmt.Println(result)
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	c.Data["json"] = responses.Response{Code: 200, Data: food}
	return
}

//@Title Detele food
//@Description Put
//@Params id path int true "food id"
//@Success 200 string
//@Failure 400 string
//@router /:id [delete]
func (c *FoodController) Delete() {
	defer c.ServeJSON()
	id, err := c.GetInt(":id")
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	c.Data["json"] = responses.Response{Code: 200, Data: id}
	return
}
