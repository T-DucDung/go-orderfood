package models

import (
	"encoding/json"
	"go-orderfood/queries"
	"log"
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

func (this *Store) GetTotalCustomer() (Store, error) {
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
