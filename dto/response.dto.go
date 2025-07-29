package dto

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid input"`
}

type UserResponse struct {
	ID        uint   `json:"id" example:"1"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Username  string `json:"username" example:"johndoe"`
}
