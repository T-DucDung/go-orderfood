package models

import (
	"encoding/json"
	"go-orderfood/queries"
)

type Shipper struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (this *Shipper) GetTotalShi(startTime, endTime int64) (string, error) {
	data, err := GetDataByQuery(queries.QueryGetCS(startTime, endTime))
	if err != nil {
		return "-1", err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return "-1", err
	}
	soluong := ""
	err = json.Unmarshal(bData, &soluong)
	if err != nil {
		return "-1", err
	}
	return soluong, nil
}
