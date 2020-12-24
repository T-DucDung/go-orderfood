package responses

import (
	"go-orderfood/models"
)

type ResShipper struct {
	Data  models.Shipper	 `json:"data"`
	Error *Err          	 `json:"error"`
}
