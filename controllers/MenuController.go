package controllers

import (
	 
	"go-orderfood/models"
	"go-orderfood/responses"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type MenuController struct {
	beego.Controller
}

//@Title Get menu
//@Description Get menu
//@Success 200 string
//@Failure 400 string
//@router / [get]
func (c *MenuController) Get() {
	defer c.ServeJSON()
	db := models.GetDatabase()
	data, err := db.Query("SELECT * FROM menu WHERE store_id = " + c.Ctx.Input.Header("StoreID"))
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	menus := make([]models.Menu, 0)
	var menu = models.Menu{}
	for data.Next() {
		err = data.Scan(&menu.CreateDate, &menu.UpdateDate, &menu.ID, &menu.Name, &menu.StoreID)
		if err != nil {
			c.Data["json"] = responses.FailResponse
			return
		}
		menus = append(menus, menu)
	}
	c.Data["json"] = responses.Response{Code: 200, Data: menus}
	return
}

//@Title Get detail menu
//@Description Get detail menu
//@Success 200 string
//@Failure 400 string
//@router /detail [get]
func (c *MenuController) GetMenu() {
	defer c.ServeJSON()
	db := models.GetDatabase()
	data, err := db.Query("SELECT * FROM menu WHERE store_id = " + c.Ctx.Input.Header("StoreID"))
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	detailmenu := make([]models.DetailMenu, 0)

	var menu = models.Menu{}
	for data.Next() {
		err = data.Scan(&menu.CreateDate, &menu.UpdateDate, &menu.ID, &menu.Name, &menu.StoreID)
		if err != nil {
			c.Data["json"] = responses.FailResponse
			return
		}
		data, err := db.Query("SELECT * FROM food WHERE menu_id = " + strconv.Itoa(menu.ID))
		if err != nil {
			fmt.Println("Menu controller 68", err)
			c.Data["json"] = responses.FailResponse
			return
		}
		foods := make([]models.Food, 0)
		var food = models.Food{}
		for data.Next() {
			err = data.Scan(&food.CreateDate, &food.UpdateDate, &food.ID, &food.Name, &food.InitPrice, &food.SalePrice, &food.Title, &food.Content, &food.StoreID, &food.MenuID, &food.CategoryID, &food.Status, &food.Image)
			if err != nil {
				fmt.Println("Menu controller 76", err)
				c.Data["json"] = responses.FailResponse
				return
			}
			foods = append(foods, food)
		}
		detailmenu = append(detailmenu, models.DetailMenu{Menu: menu,  Food: foods})
	}
	c.Data["json"] = responses.Response{Code: 200, Data: detailmenu}
	return
}

//@Title Post menu
//@Description Post
//@Params body body models.Menu true "menu"
//@Success 200 string
//@Failure 400 string
//@router / [post]
func (c *MenuController) Post() {
	defer c.ServeJSON()
	var menu models.Menu
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &menu)
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	menu.CreateDate = time.Now().Unix()
	menu.UpdateDate = time.Now().Unix()
	menu.StoreID, err = strconv.Atoi(c.Ctx.Input.Header("StoreID"))
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	db := models.GetDatabase()

	data, err := db.Prepare("INSERT INTO menu (created_date, updated_date, name, store_id) VALUES(?, ?, ?, ?);")
	if err != nil {
		panic(err.Error())
	}
	result, err := data.Exec(menu.CreateDate, menu.UpdateDate, menu.Name, menu.StoreID)
	fmt.Println(result)
	if err != nil {
		c.Data["json"] = responses.FailResponse
		return
	}
	c.Data["json"] = responses.Response{Code: 200, Data: menu}
	return
}
