package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-orderfood/queries"
	"go-orderfood/requests"
	"log"
	"strconv"
	"time"
)

type Shipper struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Image    string `json:"image"`
	UserName string `json:"user_name"`
	Status   string `json:"status"`
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
	bData, err := json.Marshal(data[0])
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

func (this *Shipper) GetTotalShipper() ([]Shipper, error) {
	data, err := GetDataByQuery("select shipper.id as id, name as name, phone as phone, username as user_name, status as status from shipper,account where shipper.id = account.typeid and type = 2")
	if err != nil {
		return []Shipper{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Shipper{}, err
	}
	shipper := []Shipper{}
	err = json.Unmarshal(bData, &shipper)
	if err != nil {
		log.Println(err)
		log.Println(shipper)
		return []Shipper{}, err
	}
	return shipper, nil
}

func (this *Shipper) CreateShipper(req requests.CreateShipper) (Shipper, error) {
	var acc = Account{
		Username: req.Username,
	}
	_, err := acc.Getaccount()
	if err == nil {
		return Shipper{}, err
	}
	time := time.Now().UnixNano() / int64(time.Millisecond)
	status := "1"

	data1, err := db.Prepare("INSERT INTO shipper(created_date, updated_date, name, phone) VALUES(?, ?, ?, ?);")
	if err != nil {
		return Shipper{}, err
	}
	_, err = data1.Exec(time, time, req.Name, req.Phone)

	num, err := GetDataByQuery("SELECT id FROM shipper ORDER BY id DESC LIMIT 1;")
	if err != nil {
		return Shipper{}, err
	}

	data, err := db.Prepare("INSERT INTO account(created_date, updated_date, username, password, type, typeid, status) VALUES(?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return Shipper{}, err
	}
	i, _ := strconv.Atoi(req.Type)
	_, err = data.Exec(time, time, req.Username, req.Password, i, num[0]["id"].(string), status)

	if err != nil {
		log.Println(err)
		return Shipper{}, err
	}
	return Shipper{}, nil
}

func (this *Shipper) DeactivateShipper(req requests.DeactivateAcc) error {

	data, err := GetDataByQuery("select count(account.typeid) from account where account.type = '" + req.Type + "' and account.typeid= '" + req.Id + "'")
	if err != nil {
		return err
	}
	if data[0]["count(account.typeid)"].(string) == "0" {
		return errors.New("User is not exist")
	}

	data1, err := db.Prepare("UPDATE account SET status = ? where type = ? and typeid = ? ")
	if err != nil {
		return err
	}

	_, err = data1.Exec(0, req.Type, req.Id)
	if err != nil {
		return err
	}
	return err
}

func (this *Shipper) ActivateShipper(req requests.DeactivateAcc) error {

	data, err := GetDataByQuery("select count(account.typeid) from account where account.type = '" + req.Type + "' and account.typeid= '" + req.Id + "'")
	if err != nil {
		return err
	}
	if data[0]["count(account.typeid)"].(string) == "0" {
		return errors.New("User is not exist")
	}

	data1, err := db.Prepare("UPDATE account SET status = ? where type = ? and typeid = ? ")
	if err != nil {
		return err
	}
	_, err = data1.Exec(1, req.Type, req.Id)
	if err != nil {
		return err
	}
	return err
}
