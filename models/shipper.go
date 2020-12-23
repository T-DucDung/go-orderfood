package models

import (
	"encoding/json"
	"fmt"
	"go-orderfood/queries"
	"log"
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

func (this *Shipper) GetTotalCustomer() (Shipper, error) {
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
