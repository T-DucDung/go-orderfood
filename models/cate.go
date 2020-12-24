package models

import (
	"encoding/json"
)

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (this *Category) GetCa() ([]Category, error) {
	data, err := GetDataByQuery("select id, name from category")
	if err != nil {
		return []Category{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Category{}, err
	}
	ls := []Category{}
	err = json.Unmarshal(bData, &ls)
	if err != nil {
		return []Category{}, err
	}
	return ls, nil
}

func (this *Category) GetRand() ([]Category, error) {
	data, err := GetDataByQuery("select id, name from category order by rand() limit 4")
	if err != nil {
		return []Category{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Category{}, err
	}
	ls := []Category{}
	err = json.Unmarshal(bData, &ls)
	if err != nil {
		return []Category{}, err
	}
	return ls, nil
}
