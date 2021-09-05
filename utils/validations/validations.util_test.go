package validations_test

import (
	"crud/api/users/dto"
	"crud/utils/validations"
	"testing"
)

func NewUserCreateDtoExample() *dto.UserCreateDto {
	return &dto.UserCreateDto{
		Name:           "Miguel",
		LastName:       "Morales",
		SecondLastName: "Reyes",
		Email:          "amorales@gmail.com",
		Password:       "admin123",
		Photo:          "sadsadasd",
		RolId:          1,
	}
}

func TestValidateSuccess(t *testing.T) {
	user := NewUserCreateDtoExample()
	err := validations.Validate(user)
	if err != nil {
		t.Error("Test don`t passed")
		t.Fail()
	}
}

func TestValidateNotHaveAProperty(t *testing.T) {
	user := NewUserCreateDtoExample()
	user.Name = ""
	err := validations.Validate(user)
	if err == nil {
		t.Error("Test don`t passed")
		t.Fail()
	}
}
