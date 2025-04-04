package api

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Apostlex0/Hotel_reservation_GO/database"
	"github.com/Apostlex0/Hotel_reservation_GO/types"
	"github.com/gofiber/fiber/v2"
)


type UserHandler struct {
	userStore db.UserStore 
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {

	var id = c.Params("id")

	// Fetch the user from the database
	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "not found"})
		}
		return err
	}

	return c.JSON(user)
}


func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	// Fetch all users from the database
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}


func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {

	var params types.CreateUserParams


	if err := c.BodyParser(&params); err != nil {
		return err
	}

	// Validate the user input
	if errors := params.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}

	// Create new user from params (includes password hashing)
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	if err := h.userStore.DeleteUser(c.Context(), userID); err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": userID})
}
func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {

	var update bson.M
	if err := c.BodyParser(&update); err != nil {
		return err
	}
	userID := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	if err := h.userStore.UpdateUser(c.Context(), filter, update); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": userID})
}
