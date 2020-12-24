package models

import (
	"encoding/json"
	"fmt"
	"go-orderfood/queries"
	"go-orderfood/requests"
	"log"
	"strconv"
	"time"
)

type Customers struct {
	ID       string `json:"id" xml:"id"`
	FullName string `json:"full_name" xml:"full_name"`
	Username string `json:"username" xml:"username"`
	Status   string `json:"status" xml:"status"`
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
	return fmt.Sprintf("%v", data[0]["soluong"]), nil
}

func (this *Customers) GetTotalCustomer() ([]Customers, error) {
	data, err := GetDataByQuery("select user.id as id, full_name as full_name, username as username, status as status from user,account where user.id = account.typeid and type = 3")
	if err != nil {
		return []Customers{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Customers{}, err
	}
	cus := []Customers{}
	err = json.Unmarshal(bData, &cus)
	if err != nil {
		log.Println(err)
		log.Println(cus)
		return []Customers{}, err
	}
	return cus, nil
}

func (this *Customers) GetInFo(id string) (Customers, error) {
	data, err := GetDataByQuery(queries.QueryGetInfoCus(id))
	if err != nil {
		return Customers{}, err
	}
	bData, err := json.Marshal(data[0])
	if err != nil {
		return Customers{}, err
	}
	c := Customers{}
	err = json.Unmarshal(bData, &c)
	if err != nil {
		return Customers{}, err
	}
	return c, err
}

func (this *Customers) CreateCustomer(req requests.CreateCustomer) (Customers, error) {
	var acc = Account{}
	_, err := acc.Getaccount()

	if err == nil {
		return Customers{}, err
	}
	time := time.Now().UnixNano() / int64(time.Millisecond)
	status := 1

	data, err := db.Prepare("INSERT INTO user(created_date, updated_date, full_name, image_url, address, mobile, email) VALUES(?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return Customers{}, err
	}
	_, err = data.Exec(time, time, req.FullName, req.ImageUrl, req.Address, req.Mobile, req.Email)

	num, err := GetDataByQuery("SELECT id FROM user ORDER BY id DESC LIMIT 1;")
	if err != nil {
		return Customers{}, err
	}

	data1, err := db.Prepare("INSERT INTO account(created_date, updated_date, username, password, type, typeid, status) VALUES(?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return Customers{}, err
	}
	i, _ := strconv.Atoi(req.Type)
	_, err = data1.Exec(strconv.FormatInt(time, 10), strconv.FormatInt(time, 10), req.UserName, req.Password, i, num[0]["id"], strconv.Itoa(status))

	if err != nil {
		return Customers{}, err
	}
	return Customers{}, nil
}
