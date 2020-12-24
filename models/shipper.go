package models

import (
	"encoding/json"
	"go-orderfood/queries"
	"log"
)

type Shipper struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
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

func (this *Shipper) GetTotalShipper() (Shipper, error){
	data, err := GetDataByQuery("select shipper.id as id, name as name, phone as phone, username as username, status as status from shipper,account where shipper.id = account.typeid and type = 2")
	if err != nil {
		return Shipper{}, err
	}
	bData, err := json.Marshal(data[0])
	if err != nil {
		return Shipper{}, err
	}
	shipper := Shipper{}
	err = json.Unmarshal(bData, &shipper)
	if err != nil {
		log.Println(err)
		log.Println(shipper)
		return Shipper{}, err
	}
	return shipper, nil
}