package models

type DetailMenu struct {
	Menu  Menu       `json:"menu" xml:"menu"`
	Foods []FoodMenu `json:"foods" xml:"foods"`
}

type Menu struct {
	CreateDate int64  `json:"create_date" xml:"create_date"`
	UpdateDate int64  `json:"update_date" xml:"update_date"`
	ID         int    `json:"id" xml:"id"`
	Name       string `json:"name" xml:"name"`
	StoreID    int    `json:"store_id" xml:"store_id"`
}
