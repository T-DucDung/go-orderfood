package models

import (
	"encoding/json"
	"orderfood/queries"
)

type Revenue struct {
	TotalRevenue string `json:"total_revenue"`
	TotalOrder   string `json:"total_order"`
}

func (this *Revenue) GetRevenue(sid int, startTime, endTime int64) (Revenue, error) {
	data, err := GetDataByQuery(queries.QueryGetR(sid, startTime, endTime))
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
