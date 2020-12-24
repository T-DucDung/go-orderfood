package models

import (
	"encoding/json"
	"go-orderfood/queries"
)

type FoodOrder struct {
	IDFood   string `json:"id_food"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
}

func (this *FoodOrder) GetListFoodOrder(id string) ([]FoodOrder, error) {
	data, err := GetDataByQuery(queries.QueryGetListFood(id))
	if err != nil {
		return []FoodOrder{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []FoodOrder{}, err
	}
	fo := []FoodOrder{}
	err = json.Unmarshal(bData, &fo)
	if err != nil {
		return []FoodOrder{}, err
	}
	return fo, nil
}
