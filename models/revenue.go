package models

import (
	"encoding/json"
	"go-orderfood/queries"
)

type Revenue struct {
	TotalRevenue string `json:"total_revenue"`
	TotalOrder   string `json:"total_order"`
}

func (this *Revenue) GetRevenueShop(sid int, startTime, endTime int64) (Revenue, error) {
	data, err := GetDataByQuery(queries.QueryGetRS(sid, startTime, endTime))
	if err != nil {
		return Revenue{}, err
	}
	if len(data) == 0 {
		return Revenue{}, nil
	}
	bData, err := json.Marshal(data[0])
	if err != nil {
		return Revenue{}, err
	}
	r := Revenue{}
	err = json.Unmarshal(bData, &r)
	if err != nil {
		return Revenue{}, err
	}
	return r, nil
}

func (this *Revenue) GetRevenue(startTime, endTime int64) (Revenue, error) {
	data, err := GetDataByQuery(queries.QueryGetR(startTime, endTime))
	if err != nil {
		return Revenue{}, err
	}
	bData, err := json.Marshal(data[0])
	if err != nil {
		return Revenue{}, err
	}
	r := Revenue{}
	err = json.Unmarshal(bData, &r)
	if err != nil {
		return Revenue{}, err
	}
	return r, nil
}
