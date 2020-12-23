package models

import (
	"encoding/json"
	"go-orderfood/queries"
	"log"
)

type Customers struct {
	ID       string `json:"id" xml:"id"`
	FullName string `json:"full_name" xml:"full_name"`
	Username string `json:"username" xml:"username"`
	Status   int    `json:"status" xml:"status"`
	ImageUrl string `xml:"image_url" json:"image_url"`
	Address  string `xml:"address" json:"address"`
	Mobile   string `xml:"mobile" json:"mobile"`
	Email    string `xml:"email" json:"email"`
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
	bData, err := json.Marshal(data[0])
	if err != nil {
		return "", err
	}
	soluong := ""
	err = json.Unmarshal(bData, &soluong)
	if err != nil {
		return "", err
	}
	return soluong, nil
}

func (this *Customers) GetTotalCustomer() (Customers, error) {
	data, err := GetDataByQuery("select user.id as id, fullname as full_name, username as username, status as status from user,account where user.id = account.typeid and type = 3")
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
