package main

import (
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(middleware.IsAuth)

	// data.Users = append(data.Users, models.User{
	// 	ID:    1,
	// 	Name:  "Andika",
	// 	Token: "abc",
	// }, models.User{
	// 	ID:    2,
	// 	Name:  "John",
	// 	Token: "def",
	// })

	// data.Todos = append(data.Todos, models.Todo{
	// 	ID:          1,
	// 	Title:       "Heloo",
	// 	IsCompleted: false,
	// 	UserID:      1,
	// })

	app.Get("/", handlers.GetAllTodosHandlers)

	app.Post("/", hanlers.CreateTodoHand)

	app.Listen(":8080")
}
