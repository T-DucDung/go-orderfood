package models

import (
	"encoding/json"
	"go-orderfood/queries"
)

type Food struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Image    string `json:"image"`
}

type FoodCate struct {
	IDCate   string `json:"id_cate"`
	NameCate string `json:"name_cate"`
	LsFood   []Food `json:"ls_food"`
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
		return Food{}, err
	}
	return cus, nil
}

func (this *Food) GetRandFood() ([]Food, error) {
	data, err := GetDataByQuery("select name as full_name, id, image from food order by rand() LIMIT 6")
	if err != nil {
		return []Food{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Food{}, err
	}
	cus := []Food{}
	err = json.Unmarshal(bData, &cus)
	if err != nil {
		return []Food{}, err
	}
	return cus, nil
}

func (this *Food) GetRandFood1(id string) ([]Food, error) {
	data, err := GetDataByQuery("select name as full_name, id, image from food where category_id = " + id + " order by rand() LIMIT 6")
	if err != nil {
		return []Food{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Food{}, err
	}
	cus := []Food{}
	err = json.Unmarshal(bData, &cus)
	if err != nil {
		return []Food{}, err
	}
	return cus, nil
}

func (this *Food) GetRandfoodCate() ([]FoodCate, error) {
	c := Category{}
	ls, err := c.GetRand()
	if err != nil {
		return []FoodCate{}, err
	}
	lsfc := []FoodCate{}
	fIns := Food{}
	for _, item := range ls {
		fc := FoodCate{}
		fc.IDCate = item.ID
		fc.NameCate = item.Name
		fc.LsFood, err = fIns.GetRandFood1(fc.IDCate)
		if err != nil {
			return []FoodCate{}, err
		}
		lsfc = append(lsfc, fc)
	}
	return lsfc, nil
}
