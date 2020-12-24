package requests

type CreateCustomer struct {
	FullName string `json:"full_name"`
	ImageUrl string `json:"image_url"`
	Address  string `json:"address"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
