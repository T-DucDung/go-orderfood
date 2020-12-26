package models

import (
	"encoding/json"
	"go-orderfood/queries"
)

type LoyalCustomers struct {
	CustomersID   string `json:"customers_id"`
	CustomersName string `json:"customers_name"`
	TotalCount    string `json:"total_count"`
}

func (this *LoyalCustomers) GetLoyalCustomers(sid int, startTime, endTime int64) ([]LoyalCustomers, error) {
	data, err := GetDataByQuery(queries.QueryGetLCN(sid, startTime, endTime))
	if err != nil {
		return []LoyalCustomers{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []LoyalCustomers{}, err
	}
	ls := []LoyalCustomers{}
	err = json.Unmarshal(bData, &ls)
	if err != nil {
		return []LoyalCustomers{}, err
	}
	return ls, nil
}
