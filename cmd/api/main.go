package main

import (
	"log"

	"go-mongo/internal/controller"
	"go-mongo/internal/database"
	"go-mongo/internal/repository"
	"go-mongo/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := database.NewMongoDatabase()
	if err != nil {
		log.Fatal("error on connecting mongodb", err)
	}

	userRepository := repository.NewUserRepository(db, "users")
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(app, userService)
	userController.Init()

	log.Fatal(app.Listen(":4000"))
}
