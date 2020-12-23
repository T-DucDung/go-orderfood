package models

import (
	"encoding/json"
	"fmt"
	"go-orderfood/queries"
	"log"
)

type Customers struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

func (this *Customers) GetCustomerName() (Customers, error) {
	data, err := GetDataByQuery(queries.QueryGetUN(this.ID))
	if err != nil {
		return Customers{}, err
	}
	bData, err := json.Marshal(data[0])
	if err != nil {
		return Customers{}, err
	}
	cus := Customers{}
	err = json.Unmarshal(bData, &cus)
	if err != nil {
		log.Println(err)
		log.Println(cus)
		return Customers{}, err
	}
	return cus, nil
}

func (this *Customers) GetTotalCus(startTime, endTime int64) (string, error) {
	data, err := GetDataByQuery(queries.QueryGetCC(startTime, endTime))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", data[0]["soluong"]), nil
}
