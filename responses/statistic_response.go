package responses

import "go-orderfood/models"

type ResStatistic struct {
	Data  models.Statistic `json:"data"`
	Error *Err             `json:"error"`
}
