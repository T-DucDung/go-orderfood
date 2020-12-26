package controllers

import (
	"go-orderfood/models"
	"go-orderfood/responses"
	"go-orderfood/services"
	"time"

	"github.com/astaxie/beego"
)

type StatisticController struct {
	beego.Controller
}

//@Title Get Statictic
//@Description Get Statictic
//@Summary "Get thống kê "
// @Params auth header string true "token"
// @Params startTime query int64 false "start time"
// @Params endTime query int64 false "end time"
//@Success 200 {object} responses.ResStatistic
//@Failure 404 responses.ResStatistic
//@router / [get]
func (this *StatisticController) GetStatistic() {
	defer this.ServeJSON()
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.ResStatisticStore{
			Data:  models.StatisticStore{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	endTime, err := this.GetInt64("endTime", time.Now().UnixNano()/int64(time.Millisecond))
	if err != nil {
		this.Data["json"] = responses.ResStatistic{
			Data:  models.Statistic{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	startTime, err := this.GetInt64("startTime", 0)
	if err != nil {
		this.Data["json"] = responses.ResStatistic{
			Data:  models.Statistic{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	s, err := services.GetStatistic(startTime, endTime)
	if err != nil {
		this.Data["json"] = responses.ResStatistic{
			Data:  models.Statistic{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	this.Data["json"] = responses.ResStatistic{
		Data:  s,
		Error: responses.NewErr(responses.Success),
	}
}
