package models

import (
	"encoding/json"
	"log"
	"orderfood/queries"
)

type Foods struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

func (this *Foods) GetFoodName() (Foods, error) {
	data, err := GetDataByQuery(queries.QueryGetFN(this.ID))
	if err != nil {
		return Foods{}, err
	}
	bData, err := json.Marshal(data[0])
	if err != nil {
		return Foods{}, err
	}
	cus := Foods{}
	err = json.Unmarshal(bData, &cus)
	if err != nil {
		log.Println(err)
		log.Println(cus)
		return Foods{}, err
	}
	return cus, nil
}
