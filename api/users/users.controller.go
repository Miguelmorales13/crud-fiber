package users

import (
	"crud/api/users/dto"
	"crud/models"
	"crud/utils/parse_body"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	s *Service
}

func NewController(app fiber.Router) (c *Controller) {
	routes := app.Group("/users")
	c = &Controller{NewService()}
	routes.Get("/", c.getAll)
	routes.Get("/:id", c.getOne)
	routes.Post("/", c.create)
	routes.Patch("/", c.update)
	routes.Delete("/:id", c.delete)
	return

}

// CreatorUsers godoc
// @Summary Create user
// @Tags users
// @Description create user
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {object} dto.UserCreateDto
// @Header 200 {string} Authorization "Bearer qwerty"
// @Failure 400,404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /users [post]
func (c *Controller) create(ctx *fiber.Ctx) error {
	user := new(dto.UserCreateDto)
	if err := parse_body.ParseBodyAndValidate(ctx, user); err != nil {
		return err
	}
	create, err := c.s.Create(user)
	if err != nil {
		return err
	}
	return ctx.JSON(models.ResponseSuccess{Data: create, Message: "Successful"})
}

// UpdateUser godoc
// @Summary Update user
// @Description update user
// @Accept  json
// @Produce  json
// @Tags users
// @Success 200 {array} dto.UserUpdateDto
// @Header 200 {string} Authorization "Bearer qwerty"
// @Failure 400,404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /users [patch]
func (c *Controller) update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	userToUpdate := new(dto.UserUpdateDto)
	if err := parse_body.ParseBodyAndValidate(ctx, userToUpdate); err != nil {
		fmt.Println(err.Error())
		return err
	}
	update, err := c.s.Update(*userToUpdate, id)
	if err != nil {
		return err
	}
	return ctx.JSON(models.ResponseSuccess{Data: update, Message: "Successful"})
}

// DeleteUser godoc
// @Summary Delete user
// @Description delete user
// @Accept  json
// @Produce  json
// @Tags users
// @Success 200 {int} 0
// @Header 200 {string} Authorization "Bearer qwerty"
// @Failure 400,404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /users [delete]
func (c *Controller) delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	rowsDeleted := c.s.Delete(id)
	return ctx.JSON(models.ResponseSuccess{Data: rowsDeleted})

}

// ListUsers godoc
// @Summary List users
// @Description get users
// @Accept  json
// @Produce  json
// @Tags users
// @Success 200 {array} dto.UserUpdateDto
// @Header 200 {string} Authorization "Bearer qwerty"
// @Failure 400,404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /users [get]
func (c *Controller) getAll(ctx *fiber.Ctx) (err error) {
	users := c.s.GetAll()
	err = ctx.JSON(models.ResponseSuccess{Data: users})
	return

}

// GetUser godoc
// @Summary Get user
// @Description get user
// @Accept  json
// @Produce  json
// @Tags users
// @Success 200 {object} dto.UserUpdateDto
// @Header 200 {string} Authorization "Bearer qwerty"
// @Failure 400,404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /users [get]
func (c *Controller) getOne(ctx *fiber.Ctx) (err error) {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	users := c.s.GetOne(id)
	err = ctx.JSON(models.ResponseSuccess{Data: users})
	return

}
