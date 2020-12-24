package responses

import (
	"go-orderfood/models"
)

type ResStore struct {
	Data  models.Store	 `json:"data"`
	Error *Err           `json:"error"`
}
