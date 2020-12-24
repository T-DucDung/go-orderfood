package requests

type CreateStore struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
