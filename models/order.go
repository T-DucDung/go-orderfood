package models

import (
	"encoding/json"
	"go-orderfood/queries"
)

type Order struct {
	ID           string      `json:"id"`
	TotalAmount  string      `json:"total_amount"`
	ListFood     []FoodOrder `json:"list_food"`
	UserID       string      `json:"user_id"`
	FullName     string      `json:"full_name"`
	Mobile       string      `json:"mobile"`
	ShipperID    string      `json:"shipper_id"`
	NameShipper  string      `json:"name_shipper"`
	PhoneShipper string      `json:"phone_shipper"`
	ShopID       string      `json:"shop_id"`
	NameShop     string      `json:"name_shop"`
	Address      string      `json:"address"`
	Status       int         `json:"status"`
}

func (this *Order) GetOrUs(id string) ([]string, error) {
	data, err := GetDataByQuery(queries.QueryGetIDOrUs(id))
	if err != nil {
		return []string{}, err
	}
	if len(data) == 0 {
		return []string{}, nil
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []string{}, err
	}
	lid := []string{}
	err = json.Unmarshal(bData, &lid)
	if err != nil {
		return []string{}, err
	}
	return lid, nil
}

func (this *Order) GetOr(id string) (Order, error) {
	data, err := GetDataByQuery(queries.QueryGetIDOrUs(id))
	if err != nil {
		return Order{}, err
	}
	if len(data) == 0 {
		return Order{}, nil
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return Order{}, err
	}
	o := Order{}
	err = json.Unmarshal(bData, &o)
	if err != nil {
		return Order{}, err
	}
	return o, nil
}
