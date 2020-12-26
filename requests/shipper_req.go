package requests

type CreateShipper struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
