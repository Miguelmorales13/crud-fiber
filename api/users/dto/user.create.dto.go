package dto

import (
	"crud/api/users/entities"
)

type UserCreateDto struct {
	Name           string `validate:"required" example:"Benito" json:"name"`
	LastName       string `validate:"required" example:"Rosas" json:"lastName"`
	Email          string `validate:"required" example:"example@gmail.com" json:"email"`
	Password       string `validate:"required" example:"sadsadasdas" json:"password"`
	SecondLastName string `validate:"required" example:"Perez" json:"secondLastName"`
	Photo          string `validate:"required" example:"http://ddsfdsf.cpom/dsfsdf.jpg" json:"Photo"`
	RolId          int    `validate:"required" example:"1" json:"rolId"`
}

func (d UserCreateDto) ToEntity() entities.User {
	return entities.User{
		Name:           d.Name,
		LastName:       d.LastName,
		Email:          d.Email,
		Password:       d.Password,
		SecondLastName: d.SecondLastName,
		Photo:          d.Photo,
		RolId:          d.RolId,
	}
}
