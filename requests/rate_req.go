package requests

type RateReq struct {
	Rate    int    `json:"rate"`
	StoreID string `json:"store_id"`
	UserID  string `json:"user_id"`
}
