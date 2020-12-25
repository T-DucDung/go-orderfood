package models

import (
	"encoding/json"
	"go-orderfood/queries"
)

type Food struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Image    string `json:"image"`
	IDShop   string `json:"id_shop"`
}

type FoodCate struct {
	IDCate   string `json:"id_cate"`
	NameCate string `json:"name_cate"`
	LsFood   []Food `json:"ls_food"`
	IDShop   string `json:"id_shop"`
}

type FoodMenu struct {
	CreateDate string `xml:"create_date" json:"create_date"`
	UpdateDate string `xml:"update_date" json:"update_date"`
	ID         string `xml:"id" json:"id"`
	Name       string `xml:"name" json:"name"`
	InitPrice  string `xml:"init_price" json:"init_price"`
	SalePrice  string `xml:"sale_price" json:"sale_price"`
	Title      string `xml:"title" json:"title"`
	Content    string `xml:"content" json:"content"`
	StoreID    string `xml:"store_id" json:"store_id"`
	MenuID     string `xml:"menu_id" json:"menu_id"`
	CategoryID string `xml:"category_id" json:"category_id"`
	Status     string `xml:"status" json:"status"`
	Image      string `xml:"image" json:"image"`
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
	data, err := GetDataByQuery("select name as full_name, id, image, store_id as id_shop from food order by rand() LIMIT 6")
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
	data, err := GetDataByQuery("select name as full_name, id, image, store_id as id_shop from food where category_id = " + id + " order by rand() LIMIT 6")
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

func (this *Food) GetFood() ([]Food, error) {
	data, err := GetDataByQuery("select name as full_name, id, image, store_id as id_shop from food")
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

func (this *Food) GetFoodByCate(cate string) ([]Food, error) {
	data, err := GetDataByQuery("select name as full_name, id, image, store_id as id_shop from food where category_id = " + cate)
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

func (this *Food) GetFoodByCateAndName(cate, name string) ([]Food, error) {
	data, err := GetDataByQuery("select name as full_name, id, image, store_id as id_shop from food where category_id = " + cate + " and name like  '%" + name + "%'")
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

func (this *Food) GetFoodByName(name string) ([]Food, error) {
	data, err := GetDataByQuery("select name as full_name, id, image, store_id as id_shop from food where name like  '%" + name + "%'")
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

func (this *Food) GetListInStore(sid string) ([]FoodMenu, error) {
	data, err := GetDataByQuery("select created_date,updated_date,id,name,init_price,sale_price,title,content,store_id,menu_id,category_id,status,image from food where store_id = " + sid)
	if err != nil {
		return []FoodMenu{}, err
	}
	bData, err := json.Marshal(data)
	if err != nil {
		return []FoodMenu{}, err
	}
	cus := []FoodMenu{}
	err = json.Unmarshal(bData, &cus)
	if err != nil {
		return []FoodMenu{}, err
	}
	return cus, nil
}
