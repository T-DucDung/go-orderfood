package controllers

import (
	"go-orderfood/models"
	"go-orderfood/responses"
	"go-orderfood/services"
	"time"

	"github.com/astaxie/beego"
)

type StatisticStoreController struct {
	beego.Controller
}

//@Title Get Statictic Store
//@Description Get Statictic Store
//@Summary "Get thống kê một cửa hàng"
// @Params auth header string true "store id"
// @Params sid query int true "store id"
// @Params startTime query int64 false "start time"
// @Params endTime query int64 false "end time"
//@Success 200 {object} responses.ResStatisticStore
//@Failure 404 responses.ResStatisticStore
//@router / [get]
func (this *StatisticStoreController) GetStatisticStore() {
	defer this.ServeJSON()
	sid, err := this.GetInt("sid", -1)
	if err != nil {
		this.Data["json"] = responses.ResStatisticStore{
			Data:  models.StatisticStore{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	endTime, err := this.GetInt64("endTime", time.Now().UnixNano()/int64(time.Millisecond))
	if err != nil {
		this.Data["json"] = responses.ResStatisticStore{
			Data:  models.StatisticStore{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	startTime, err := this.GetInt64("startTime", 0)
	if err != nil {
		this.Data["json"] = responses.ResStatisticStore{
			Data:  models.StatisticStore{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	sStore, err := services.GetStatisticStore(sid, startTime, endTime)
	if err != nil {
		this.Data["json"] = responses.ResStatisticStore{
			Data:  models.StatisticStore{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	this.Data["json"] = responses.ResStatisticStore{
		Data:  sStore,
		Error: responses.NewErr(responses.Success),
	}
}
