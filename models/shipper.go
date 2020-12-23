package models

import (
	"encoding/json"
	"fmt"
	"go-orderfood/queries"
)

type Shipper struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (this *Shipper) GetTotalShi(startTime, endTime int64) (string, error) {
	data, err := GetDataByQuery(queries.QueryGetCS(startTime, endTime))
	if err != nil {
		return "-1", err
	}
	return fmt.Sprintf("%v", data[0]["soluong"]), nil
}

func (this *Shipper) GetInfoShipper(id string) (Shipper, error) {
	data, err := GetDataByQuery(queries.QueryGetShip(id))
	if err != nil {
		return Shipper{}, err
	}
	if len(data) == 0 {
		return Shipper{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return Shipper{}, err
	}
	s := Shipper{}
	err = json.Unmarshal(bData, &s)
	if err != nil {
		return Shipper{}, err
	}
	return s, nil
}
