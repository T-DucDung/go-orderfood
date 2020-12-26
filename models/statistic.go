package models

type Statistic struct {
	TotalShop         string           `json:"total_shop"`
	TotalCustomer     string           `json:"total_customer"`
	TotalShipper      string           `json:"total_shipper"`
	TotalOrder        string           `json:"total_order"`
	Revenue           string           `json:"revenue"`
	StatisticAllStore []StatisticStore `json:"statistic_all_store"`
}
