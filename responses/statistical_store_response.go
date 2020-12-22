package responses

import (
	"orderfood/models"
)

type ResStatisticStore struct {
	Data  models.StatisticStore `json:"data"`
	Error *Err                  `json:"error"`
}
