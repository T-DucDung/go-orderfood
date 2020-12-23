package models

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Typeid   string `json:"typeid"`
}
type Authorize struct {
	Username string `xml:"username" json:"username"`
	Password string `xml:"password" json:"password"`
}

func (this *Account) Getaccount() (Account, error) {
	data, err := GetDataByQuery("select username,password, type,typeid from account where username='" + this.Username + "'")
	if err != nil {
		return Account{}, err
	}
	if len(data) == 0 {
		return Account{}, errors.New("User's not exist")
	}
	bData, err := json.Marshal(data[0])
	if err != nil {
		return Account{}, err
	}
	acc := Account{}
	err = json.Unmarshal(bData, &acc)
	if err != nil {
		return Account{}, err
	}
	return acc, nil
}


func (this *Account) CreateAcc() (Account, error) {
	_, err := this.Getaccount()
	var store = Store{}
	var shipper = Shipper{}
	var customer = Customers{}

	if err == nil {
		return Account{}, err
	}
	time := time.Now().Unix()
	status := 1
	if this.Type == "1" {
		num, err := GetDataByQuery("SELECT id FROM store ORDER BY id DESC LIMIT 1;")
		if err != nil {
			return Account{}, err
		}
		data1, err := db.Prepare("INSERT INTO store(created_date, updated_date, name, rate_avg, rate_one, rate_two, rate_three, rate_four, rate_five) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			return Account{}, err
		}
		_, err = data1.Exec(strconv.FormatInt(time, 10), strconv.FormatInt(time, 10), store.Name, 0, 0, 0, 0, 0)
		
		data, err := db.Prepare("INSERT INTO account(created_date, updated_date, username, password, type, typeid, status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			return Account{}, err
		}
		_, err = data.Exec(strconv.FormatInt(time, 10), strconv.FormatInt(time, 10), this.Username ,this.Password ,this.Type ,num[0]["id"] ,strconv.Itoa(status))
	} else if this.Type == "2" {
		num, err := GetDataByQuery("SELECT id FROM shipper ORDER BY id DESC LIMIT 1;")
		if err != nil {
			return Account{}, err
		}
		data1, err := db.Prepare("INSERT INTO shipper(created_date, updated_date, name, phone) VALUES(?, ?, ?, ?);")
		if err != nil {
			return Account{}, err
		}
		_, err = data1.Exec(strconv.FormatInt(time, 10), strconv.FormatInt(time, 10), shipper.Name, shipper.Phone)
		
		data, err := db.Prepare("INSERT INTO account(created_date, updated_date, username, password, type, typeid, status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			return Account{}, err
		}
		_, err = data.Exec(strconv.FormatInt(time, 10), strconv.FormatInt(time, 10), this.Username ,this.Password ,this.Type ,num[0]["id"] ,strconv.Itoa(status))
	}else if this.Type == "3" {
		num, err := GetDataByQuery("SELECT id FROM user ORDER BY id DESC LIMIT 1;")
		if err != nil {
			return Account{}, err
		}
		data1, err := db.Prepare("INSERT INTO user(created_date, updated_date, full_name, image_url, address, mobile, email) VALUES(?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			return Account{}, err
		}
		_, err = data1.Exec(strconv.FormatInt(time, 10), strconv.FormatInt(time, 10), customer.FullName, customer.ImageUrl, customer.Address, customer.Mobile, customer.Email)
		
		data1, err = db.Prepare("INSERT INTO account(created_date, updated_date, username, password, type, typeid, status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			return Account{}, err
		}
		_, err = data1.Exec(strconv.FormatInt(time, 10), strconv.FormatInt(time, 10), this.Username ,this.Password ,this.Type ,num[0]["id"] ,strconv.Itoa(status))
	}else {
		data, err := db.Prepare("INSERT INTO account(created_date, updated_date, username, password, type, typeid, status) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			return Account{}, err
		}
		_, err = data.Exec(strconv.FormatInt(time, 10), strconv.FormatInt(time, 10), this.Username ,this.Password ,this.Type ,0 ,strconv.Itoa(status))
	}

		
	if err != nil {
		return Account{}, err
	}
	return Account{}, nil
}
