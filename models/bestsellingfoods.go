package models

import (
	"encoding/json"
	"go-orderfood/queries"
)

type BestSellingFoods struct {
	FoodID     string `json:"food_id"`
	FoodName   string `json:"food_name"`
	TotalCount string `json:"total_count"`
}

func (this *BestSellingFoods) GetBestSellingFoods(sid int, startTime, endTime int64) ([]BestSellingFoods, error) {
	data, err := GetDataByQuery(queries.QueryGetBSF(sid, startTime, endTime))
	if err != nil {
		return []BestSellingFoods{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []BestSellingFoods{}, err
	}
	ls := []BestSellingFoods{}
	err = json.Unmarshal(bData, &ls)
	if err != nil {
		return []BestSellingFoods{}, err
	}
	return ls, nil
}
