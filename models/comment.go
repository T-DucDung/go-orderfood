package models

type Comment struct {
	CreateDate int64  `json:"create_date" xml:"create_date"`
	UpdateDate int64  `json:"update_date" xml:"update_date"`
	ID         int    `json:"id" xml:"id"`
	StoreID    int    `json:"store_id" xml:"store_id"`
	Comment    string `json:"comment" xml:"comment"`
	UserID     int    `json:"user_id" xml:"user_id"`
}
