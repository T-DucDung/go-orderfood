package models

import (
	"encoding/json"
	"fmt"
	"go-orderfood/queries"
	"go-orderfood/requests"
	"strconv"
	"time"
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

func (this *Order) CreateOrder(o requests.ReqOrder, sid string) error {
	t := time.Now().UnixNano() / int64(time.Millisecond)
	data, err := db.Prepare("INSERT INTO `order` (created_date,updated_date,order_status,user_id,address,store_id,shipper_id, options_orderfood) values(?,?,?,?,?,?,?,?);")
	if err != nil {
		return err
	}
	_, err = data.Exec(t, t, 1, sid, o.Address, o.StoreID, 0, o.OptionsOrderfood)
	if err != nil {
		return err
	}
	d, err := GetDataByQuery("select id from `order` order by id desc limit 1")
	if err != nil {
		return err
	}
	total := 0
	for _, item := range o.ListFood {
		data, err = db.Prepare("INSERT INTO orderfood (created_date,updated_date,order_id,food_id,quantity,price) values(?,?,?,?,?,?)")
		if err != nil {
			return err
		}
		f := Food{}
		p, err := f.Price(item.IDFood, item.Quantity)
		if err != nil {
			return err
		}
		total = total + p
		_, err = data.Exec(t, t, d[0]["id"], item.IDFood, item.Quantity, p)
		if err != nil {
			return err
		}
	}
	return ExecNonQuery("update `order` set total_amount = " + strconv.Itoa(total) + " , order_status = 2 where id = " + d[0]["id"].(string))
}
