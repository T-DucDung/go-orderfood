package responses

import (
	"go-orderfood/models"
)

type ResListOrder struct {
	Data  []models.Order `json:"data"`
	Error *Err           `json:"error"`
}
