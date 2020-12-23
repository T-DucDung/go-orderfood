package models

import (
	"encoding/json"
	"go-orderfood/queries"
	"log"
)

type Food struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

func (this *Food) GetFoodName() (Food, error) {
	data, err := GetDataByQuery(queries.QueryGetFN(this.ID))
	if err != nil {
		return Food{}, err
	}
	bData, err := json.Marshal(data[0])
	if err != nil {
		return Food{}, err
	}
	cus := Food{}
	err = json.Unmarshal(bData, &cus)
	if err != nil {
		log.Println(err)
		log.Println(cus)
		return Food{}, err
	}
	return cus, nil
}
