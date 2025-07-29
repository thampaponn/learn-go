package dto

type SignUpInput struct {
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Username  string `json:"username" example:"johndoe"`
	Password  string `json:"password" example:"password123"`
}

type LoginInput struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"securepassword123"`
}
