package models

type Category struct {
	CreateDate int64  `json:"create_date" xml:"create_date"`
	UpdateDate int64  `json:"update_date" xml:"update_date"`
	ID         int    `json:"id" xml:"id"`
	Name       string `json:"name" xml:"name"`
	IsActive   bool   `json:"is_active" xml:"is_active"`
}
