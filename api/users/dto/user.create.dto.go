package dto

import (
	"crud/api/users/entities"
)

type UserCreateDto struct {
	Name           string `validate:"required" example:"Benito" json:"name"`
	LastName       string `example:"Rosas" json:"lastName"`
	Email          string `example:"example@gmail.com" json:"email"`
	Password       string `example:"sadsadasdas" json:"password"`
	SecondLastName string `example:"Perez" json:"secondLastName"`
	Photo          string `example:"http://ddsfdsf.cpom/dsfsdf.jpg" json:"Photo"`
	RolId          int    `example:"1" json:"rolId"`
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
