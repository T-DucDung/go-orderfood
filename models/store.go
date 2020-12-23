package models

import (
	"encoding/json"
	"go-orderfood/queries"
)

type Store struct {
	ID          string `json:"id"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	Name        string `json:"name"`
	RateAvg     string `json:"rate_avg"`
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
