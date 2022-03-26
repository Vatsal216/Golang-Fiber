package handlers

import "github.com/gofiber/fiber"

func GetAllTodosHandlers(c *fiber.Ctx) error {
	todos := []models.Todo{}
	userid := c.Locals("userid").(uint)

	for _, t := range data.Todos {
		if t.UserID == userid {
			todos = append(todos, t)
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": todos,
	})
}
