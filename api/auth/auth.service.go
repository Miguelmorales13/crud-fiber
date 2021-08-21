package auth

import (
	"crud/api/auth/dto"
	"crud/api/users"
	"crud/api/users/entities"
	"crud/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

var service *Service

type Service struct {
	userService *users.Service
}

func NewService() *Service {
	if service == nil {
		service = &Service{users.NewService()}
	}
	return service
}
func (s *Service) signIn(signIn dto.SignInRequest) (user entities.User, token string, err error) {
	userWanted := s.userService.GetOneByEmail(signIn.User)
	if userWanted.ID == 0 {
		err = fiber.ErrUnauthorized
		return
	}
	if isValid := userWanted.IsValidPassword(signIn.Password); !isValid {
		err = fiber.ErrUnauthorized
		return
	}
	token = jwt.Generate(&jwt.TokenPayload{
		ID: userWanted.ID,
	})
	user = userWanted
	return
}
