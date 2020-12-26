package responses

import (
	"go-orderfood/models"
)

type ResCustomer struct {
	Data  []models.Customers `json:"data"`
	Error *Err               `json:"error"`
}
