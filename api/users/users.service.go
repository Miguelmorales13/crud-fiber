package users

import (
	"crud/api/users/dto"
	entities "crud/api/users/entities"
	"crud/providers"
	"github.com/gofiber/fiber/v2"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
func (s *Service) Create(user *dto.UserCreateDto) (userCreated entities.User, e error) {
	userCreated = user.ToEntity()
	create := providers.Database.Create(&userCreated)
	return userCreated, create.Error
}

func (s *Service) Update(userUpdated dto.UserUpdateDto, id int) (*entities.User, error) {
	user := s.GetOne(id)
	if user.ID == 0 {
		return &entities.User{}, fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	updates := providers.Database.Model(&user).Updates(userUpdated)
	return &user, updates.Error
}
func (s *Service) Delete(id int) int64 {
	user := s.GetOne(id)
	if user.ID == 0 {
		return 0
	}
	deleted := providers.Database.Delete(&user)
	return deleted.RowsAffected
}
func (s *Service) GetAll() (users []entities.User) {
	providers.Database.Find(&users)
	return
}

func (s *Service) GetOne(id int) (user entities.User) {
	providers.Database.Find(&user, id)
	return
}

func (s *Service) GetOneByEmail(email string) (user entities.User) {
	providers.Database.Where(map[string]interface{}{"email": email}).First(&user)
	return
}
