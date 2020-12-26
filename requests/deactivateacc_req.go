package requests

type DeactivateAcc struct {
	Id 		 string `json:"id"`
	// Username string `json:"username"`
	Type     string `json:"type"`
}