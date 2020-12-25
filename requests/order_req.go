package requests

type ReqOrder struct {
	ID        string  `json:"id"`
	Address   string  `json:"address"`
	ShipperID string  `json:"shipper_id"`
	StoreID   string  `json:"store_id"`
	ListFood  []Foods `json:"list_food"`
}

type Foods struct {
	IDFood   string `json:"id_food"`
	Quantity string `json:"quantity"`
}
