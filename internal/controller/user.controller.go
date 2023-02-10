package controller

import (
	"fmt"

	"go-mongo/internal/contract"
	"go-mongo/internal/model"
	"go-mongo/internal/service"
	"go-mongo/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Init()
	InsertOne(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	FindOne(ctx *fiber.Ctx) error
	UpdateOne(ctx *fiber.Ctx) error
	DeleteOne(ctx *fiber.Ctx) error
}

type userController struct {
	app *fiber.App
	s   service.UserService
}

func NewUserController(app *fiber.App, s service.UserService) UserController {
	return &userController{
		app: app,
		s:   s,
	}
}

func (c *userController) Init() {
	c.app.Post("/users", c.InsertOne)
	c.app.Get("/users/:id", c.FindOne)
}

func (c *userController) InsertOne(ctx *fiber.Ctx) error {
	var body contract.InsertOneUserReqBody
	if err := ctx.BodyParser(&body); err != nil {
		ctx.Status(fiber.StatusUnprocessableEntity).JSON(contract.ErrorResponse{
			OK:         false,
			StatusCode: fiber.StatusUnprocessableEntity,
			Error:      err.Error(),
		})
		return nil
	}
	if errs, err := validator.ValidateStruct(&body); err != nil {
		ctx.Status(fiber.StatusUnprocessableEntity).JSON(contract.ErrorResponse{
			OK:                false,
			StatusCode:        fiber.StatusUnprocessableEntity,
			Error:             err.Error(),
			StructErrorFields: errs,
		})
		return nil
	}

	user := model.User{
		Name:     body.Name,
		Username: body.Username,
		Password: body.Password,
	}

	res, err := c.s.InsertOne(ctx.Context(), user)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusCreated).JSON(contract.Response{
		OK:         true,
		StatusCode: fiber.StatusCreated,
		Data:       res,
	})

	return nil
}

func (c *userController) Find(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

func (c *userController) FindOne(ctx *fiber.Ctx) error {
	var params contract.FindOneUserReqParams
	if err := ctx.ParamsParser(&params); err != nil {
		return err
	}

	fmt.Println(params)

	return nil
}

func (c *userController) UpdateOne(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

func (c *userController) DeleteOne(ctx *fiber.Ctx) error {
	panic("unimplemented")
}
