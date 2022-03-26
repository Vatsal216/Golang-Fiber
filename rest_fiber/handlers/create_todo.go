package handlers

import "github.com/gofiber/fiber"

func CreateTodoHand(c *fiber.Ctx) error {
	var input models.CreateTodoRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})

	}

	data.Todos = append(data.Todos, models.Todo{
		ID:          2,
		Title:       input.Title,
		IsCompleted: false,
		UserID:      1,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   nil,
		"message": "todo created",
	})
}
