package models

import (
	"encoding/json"
	"go-orderfood/queries"
	"go-orderfood/requests"
	"log"
	"strconv"
	"time"
)

type Store struct {
	CreateDate string `json:"create_date" xml:"create_date"`
	UpdateDate string `json:"update_date" xml:"update_date"`
	ID         string `json:"id" xml:"id"`
	Name       string `json:"name" xml:"name"`
	RateAvg    string `json:"rate_avg" xml:"rate_avg"`
	RateOne    string `json:"rate_one" xml:"rate_one"`
	RateTwo    string `json:"rate_two" xml:"rate_two"`
	RateThree  string `json:"rate_three" xml:"rate_three"`
	RateFour   string `json:"rate_four" xml:"rate_four"`
	RateFive   string `json:"rate_five" xml:"rate_five"`
	Username   string `json:"username"`
	Status     string `json:"status"`
}

func (this *Store) GetStoreInfo() (Store, error) {
	data, err := GetDataByQuery(queries.QueryGetIS(this.ID))
	if err != nil {
		return Store{}, err
	}
	if len(data) == 0 {
		return Store{}, err
	}
	bData, err := json.Marshal(data[0])
	if err != nil {
		return Store{}, err
	}
	store := Store{}
	err = json.Unmarshal(bData, &store)
	if err != nil {
		return Store{}, err
	}
	return store, nil
}

func (this *Store) GetID(startTime, endTime int64) ([]LString, error) {
	data, err := GetDataByQuery(queries.QueryGetIdS(startTime, endTime))
	if err != nil {
		return []LString{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []LString{}, err
	}
	ls := []LString{}
	err = json.Unmarshal(bData, &ls)
	if err != nil {
		return []LString{}, err
	}
	return ls, err
}

type LString struct {
	ID string `json:"id" xml:"id"`
}

func (this *Store) GetTotalStore() ([]Store, error) {
	data, err := GetDataByQuery("select store.id as id, name as name, username as username, status as status from store,account where store.id = account.typeid and type=1")
	if err != nil {
		return []Store{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Store{}, err
	}
	store := []Store{}
	err = json.Unmarshal(bData, &store)
	if err != nil {
		log.Println(err)
		log.Println(store)
		return []Store{}, err
	}
	return store, nil
}

func (this *Store) CreateStore(req requests.CreateStore) (Store, error) {
	var acc = Account{
		Username: req.Username,
	}
	_, err := acc.Getaccount()
	if err == nil {
		return Store{}, err
	}
	time := time.Now().UnixNano() / int64(time.Millisecond)
	status := "1"

	data1, err := db.Prepare("INSERT INTO store(created_date, updated_date, name, rate_avg, rate_one, rate_two, rate_three, rate_four, rate_five) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return Store{}, err
	}
	_, err = data1.Exec(time, time, req.Name, 0, 0, 0, 0, 0, 0)

	num, err := GetDataByQuery("SELECT id FROM store ORDER BY id DESC LIMIT 1;")
	if err != nil {
		return Store{}, err
	}

	data, err := db.Prepare("INSERT INTO account(created_date, updated_date, username, password, type, typeid, status) VALUES(?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return Store{}, err
	}
	i, _ := strconv.Atoi(req.Type)
	_, err = data.Exec(time, time, req.Username, req.Password, i, num[0]["id"].(string), status)

	if err != nil {
		log.Println("err", err)
		return Store{}, err
	}
	return Store{}, nil
}
