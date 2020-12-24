package requests

type CreateStore struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
