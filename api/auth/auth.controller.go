package auth

import (
	"crud/api/auth/dto"
	"crud/models"
	"crud/utils/parse_body"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	s *Service
}

func NewController(app fiber.Router) (c *Controller) {
	routes := app.Group("/auth")
	c = &Controller{NewService()}
	routes.Post("/sign-in", c.signIn)
	return

}

// SignIn godoc
// @Summary Sign in
// @Tags auth
// @Description sign in
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.SignInRequest
// @Failure 400,404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /auth/sign-in [post]
func (c *Controller) signIn(ctx *fiber.Ctx) error {
	signIn := new(dto.SignInRequest)

	if err := parse_body.ParseBodyAndValidate(ctx, signIn); err != nil {
		return err
	}
	user, token, err := c.s.signIn(*signIn)
	if err != nil {
		return err
	}
	return ctx.JSON(models.ResponseSuccess{Data: dto.SignInResponse{user, token}, Message: "login successful"})
}
