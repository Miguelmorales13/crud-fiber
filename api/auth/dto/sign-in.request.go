package dto

type SignInRequest struct {
	User     string `validate:"required" example:"example@example.com" json:"user"`
	Password string `validate:"required" example:"******" json:"password"`
}
