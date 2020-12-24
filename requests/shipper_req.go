package requests

type CreateShipper struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
