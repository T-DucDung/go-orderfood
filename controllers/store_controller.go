package controllers

import (
	"encoding/json"
	"fmt"
	"go-orderfood/models"
	"go-orderfood/requests"
	"go-orderfood/responses"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

type StoreController struct {
	beego.Controller
}

//@Title Post Comment
//@Description Post Comment
//@Summary "Post Comment"
// @Params auth header string true "token"
// @Param data body models.Comment true "comment"
//@Success 200 responses.Response
//@Failure 404 responses.Response
//@router /comment [post]
func (this *StoreController) CreateComment() {
	defer this.ServeJSON()
	sid, err := strconv.Atoi(this.Ctx.Request.Header.Get("Id"))
	if err != nil {
		this.Data["json"] = responses.ResStatisticStore{
			Data:  models.StatisticStore{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "3" {
		this.Data["json"] = responses.ResStatisticStore{
			Data:  models.StatisticStore{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	log.Println(sid)

	c := models.Comment{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &this)
	c.IDUser = strconv.Itoa(sid)
	log.Println(this)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	err = c.CreateCom()
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 201, Message: "Error"}}
		return
	}
	this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 200, Message: "Success"}}
}

//@Title Get Comment
//@Description Get Comment
//@Summary "Get Comment"
// @Params sid query string true "sid"
//@Success 200 responses.Response
//@Failure 404 responses.Response
//@router /comments [get]
func (this *StoreController) Comment() {
	defer this.ServeJSON()
	db := models.GetDatabase()
	sid := this.GetString("sid")
	data, err := db.Query("SELECT created_date, updated_date, id, store_id, comment, user_id FROM comment WHERE store_id = " + sid)
	if err != nil {
		fmt.Println("store controller", err)
		this.Data["json"] = responses.FailResponse
		return
	}
	comments := make([]models.Comment, 0)
	var comment = models.Comment{}
	for data.Next() {
		err = data.Scan(&comment.CreateDate, &comment.UpdateDate, &comment.ID, &comment.IDStore, &comment.Comment, &comment.IDUser)
		if err != nil {
			fmt.Println("store controller", err)
			this.Data["json"] = responses.FailResponse
			return
		}
		comments = append(comments, comment)
	}
	this.Data["json"] = responses.Response{Error: &responses.Err{Code: 200, Message: "Success"}, Data: comments}
	return
}

//@Title Set Rate
//@Description Set Rate
//@Summary "Set Rate"
// @Params auth header string true "token"
// @Params data body requests.RateReq true "sid"
//@Success 200 responses.Response
//@Failure 404 responses.Response
//@router /rate [post]
func (this *StoreController) UpdateRate() {
	defer this.ServeJSON()
	sid, err := strconv.Atoi(this.Ctx.Request.Header.Get("Id"))
	if err != nil {
		this.Data["json"] = responses.Response{
			Data:  "",
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "3" {
		this.Data["json"] = responses.Response{
			Data:  "",
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	log.Println(sid)
	req := requests.RateReq{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Data["json"] = responses.Response{
			Data:  "",
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	req.UserID = strconv.Itoa(sid)
	s := models.Store{}
	err = s.SetRate(req)
	if err != nil {
		this.Data["json"] = responses.Response{
			Data:  "",
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responses.Response{
		Data:  "",
		Error: responses.NewErr(responses.Success),
	}
	return
}

//@Title Get Rate
//@Description Get Rate
//@Summary "Get Rate"
// @Params sid query string true "sid"
//@Success 200 responses.Response
//@Failure 404 responses.Response
//@router /rates [get]
func (this *StoreController) Rate() {
	defer this.ServeJSON()
	sid := this.GetString("sid")
	s := models.Store{}
	rate, err := s.GetRate(sid)
	if err != nil {
		this.Data["json"] = responses.Response{Error: &responses.Err{Code: 201, Message: "UnSuccess"}, Data: "rate"}
		return
	}
	this.Data["json"] = responses.Response{Error: &responses.Err{Code: 200, Message: "Success"}, Data: rate}
	return
}
