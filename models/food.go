package models

import (
	"encoding/json"
	"go-orderfood/queries"
	"log"
	"strconv"
)

type Food struct {
	CreateDate int64  `xml:"create_date" json:"create_date"`
	UpdateDate int64  `xml:"update_date" json:"update_date"`
	ID         int    `xml:"id" json:"id"`
	Name       string `xml:"name" json:"name"`
	InitPrice  string `xml:"init_price" json:"init_price"`
	SalePrice  string `xml:"sale_price" json:"sale_price"`
	Title      string `xml:"title" json:"title"`
	Content    string `xml:"content" json:"content"`
	StoreID    int    `xml:"store_id" json:"store_id"`
	MenuID     int    `xml:"menu_id" json:"menu_id"`
	CategoryID int    `xml:"category_id" json:"category_id"`
	Status     int    `xml:"status" json:"status"`
	Image      string `xml:"image" json:"image"`
}

func (this *Food) GetFoodName() (Food, error) {
	data, err := GetDataByQuery(queries.QueryGetFN(strconv.Itoa(this.ID)))
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
