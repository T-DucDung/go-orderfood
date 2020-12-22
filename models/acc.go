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
	if err == nil {
		return Account{}, err
	}
	time := time.Now().Unix()
	status := 1
	_, err = ExecNonQuery("insert into account(created_date, updated_date, username, password, type, typeid, status) values(" + strconv.FormatInt(time, 10) + "," + strconv.FormatInt(time, 10) + ",'" + this.Username + "','" + this.Password + "'," + this.Type + "," + this.Typeid + "," + strconv.Itoa(status) + ")")
}
