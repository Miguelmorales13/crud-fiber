package dto

import "crud/api/users/entities"

type SignInResponse struct {
	User  entities.User ` json:"user"`
	Token string        `json:"token"`
}
