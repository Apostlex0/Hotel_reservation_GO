package api

import (
	"github.com/Apostlex0/Hotel_reservation_GO/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName: "Bond",
	}
	return c.JSON(u)
}
func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"name": "John Do e"})
}
