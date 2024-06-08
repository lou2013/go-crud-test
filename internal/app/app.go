package app

import (
	"errors"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	userService "test.com/test/internal/app/service"
)

func Start() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	api := app.Group("/api")

	api.Get("/users", func(c *fiber.Ctx) error {
		users := userService.GetAll()

		return c.Status(fiber.StatusOK).JSON(users)
	})

	api.Get("/users/:userId", func(c *fiber.Ctx) error {
		uId, err := strconv.Atoi(c.Params("userId"))
		userId := int32(uId)
		if err != nil {
			return err
		}
		user, err := userService.GetById(userId)

		if err != nil {
			if errors.Is(err, fiber.ErrNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Resource not found",
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(user)
	})

	api.Put("/users/:userId", func(c *fiber.Ctx) error {
		uId, err := strconv.Atoi(c.Params("userId"))
		userId := int32(uId)
		var u userService.User
		json.Unmarshal(c.Body(), &u)

		if err != nil {
			return err
		}
		newUser, err := userService.Update(userId, u)

		if err != nil {
			if errors.Is(err, fiber.ErrNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Resource not found",
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(newUser)
	})

	api.Post("/users", func(c *fiber.Ctx) error {
		var u userService.User
		err := json.Unmarshal(c.Body(), &u)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		newUser := userService.Add(u)

		return c.Status(fiber.StatusOK).JSON(newUser)
	})

	api.Delete("/users/:userId", func(c *fiber.Ctx) error {
		uId, err := strconv.Atoi(c.Params("userId"))
		userId := int32(uId)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		deletedUser, err := userService.Remove(userId)

		if err != nil {
			if errors.Is(err, fiber.ErrNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Resource not found",
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(deletedUser)
	})

	app.Listen(":3005")
}
