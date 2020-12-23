package models

import (
	"encoding/json"
	"go-orderfood/queries"
	"strconv"
	"log"
)

type Store struct {
	CreateDate int64   `json:"create_date" xml:"create_date"`
	UpdateDate int64   `json:"update_date" xml:"update_date"`
	ID         int     `json:"id" xml:"id"`
	Name       string  `json:"name" xml:"name"`
	RateAvg    float32 `json:"rate_avg" xml:"rate_avg"`
	RateOne    float32 `json:"rate_one" xml:"rate_one"`
	RateTwo    float32 `json:"rate_two" xml:"rate_two"`
	RateThree  float32 `json:"rate_three" xml:"rate_three"`
	RateFour   float32 `json:"rate_four" xml:"rate_four"`
	RateFive   float32 `json:"rate_five" xml:"rate_five"`
	Username   string  `json:"username"`
	Status     int 	   `json:"status"`
}

func (this *Store) GetStoreInfo() (Store, error) {
	data, err := GetDataByQuery(queries.QueryGetIS(strconv.Itoa(this.ID)))
	if err != nil {
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

func (this *Store) GetID(startTime, endTime int64) ([]string, error) {
	data, err := GetDataByQuery(queries.QueryGetIdS(startTime, endTime))
	if err != nil {
		return []string{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []string{}, err
	}
	ls := []string{}
	err = json.Unmarshal(bData, &ls)
	if err != nil {
		return []string{}, err
	}
	return ls, err
}

func (this *Store) GetTotalCustomer() (Store, error){
	data, err := GetDataByQuery("select store.id as id, name as name, username as username, status as status, rate_avg as rate_avg from store,account where store.id = account.typeid and type=1")
	if err != nil {
		return Store{}, err
	}
	bData, err := json.Marshal(data[0])
	if err != nil {
		return Store{}, err
	}
	store := Store{}
	err = json.Unmarshal(bData, &store)
	if err != nil {
		log.Println(err)
		log.Println(store)
		return Store{}, err
	}
	return store, nil
}