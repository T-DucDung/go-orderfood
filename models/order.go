package models

import (
	"encoding/json"
	"fmt"
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
	OrderStatus  string      `json:"order_status"`
}

func (this *Order) GetOrUs(id string) ([]Order, error) {
	data, err := GetDataByQuery(queries.QueryGetIDOrUs(id))
	if err != nil {
		return []Order{}, err
	}
	if len(data) == 0 {
		return []Order{}, nil
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Order{}, err
	}
	lid := []Order{}
	err = json.Unmarshal(bData, &lid)
	if err != nil {
		return []Order{}, err
	}
	return lid, nil
}

func (this *Order) GetOrSh() ([]Order, error) {
	data, err := GetDataByQuery(queries.QueryGetOrSh())
	if err != nil {
		return []Order{}, err
	}
	if len(data) == 0 {
		return []Order{}, nil
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Order{}, err
	}
	lid := []Order{}
	err = json.Unmarshal(bData, &lid)
	if err != nil {
		return []Order{}, err
	}
	return lid, nil
}

func (this *Order) GetOrS(id string) ([]Order, error) {
	data, err := GetDataByQuery(queries.QueryGetOrS(id))
	if err != nil {
		return []Order{}, err
	}
	if len(data) == 0 {
		return []Order{}, nil
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Order{}, err
	}
	lid := []Order{}
	err = json.Unmarshal(bData, &lid)
	if err != nil {
		return []Order{}, err
	}
	return lid, nil
}

func (this *Order) GetOrSh1(id string) ([]Order, error) {
	data, err := GetDataByQuery(queries.QueryGetOrSh1(id))
	if err != nil {
		return []Order{}, err
	}
	if len(data) == 0 {
		return []Order{}, nil
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []Order{}, err
	}
	lid := []Order{}
	err = json.Unmarshal(bData, &lid)
	if err != nil {
		return []Order{}, err
	}
	return lid, nil
}

func (this *Order) GetStatus(id string) (string, error) {
	data, err := GetDataByQuery(queries.QueryGetStatus(id))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", data[0]["order_status"]), nil
}

func (this *Order) GetStoreOrStatus(id string) (string, string, error) {
	data, err := GetDataByQuery(queries.QueryGetStoreOrStaus(id))
	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf("%v", data[0]["order_status"]), fmt.Sprintf("%v", data[0]["store_id"]), nil
}

func (this *Order) UpStatus(ido, ids string) error {
	return ExecNonQuery(queries.QueryUpStatus(ido, ids))
}

func (this *Order) EndStatus(ido, ids string) error {
	return ExecNonQuery(queries.QueryEndStatus(ido, ids))
}

func (this *Order) UpStatusStore(id string) error {
	return ExecNonQuery(queries.QueryUpStoreOrStatus(id))
}

// func (this *Order) CreateOrder() error {

// }
