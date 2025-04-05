package api

import (
	"github.com/Apostlex0/Hotel_reservation_GO/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// It processes requests for listing hotels, getting hotel details, and viewing rooms
type HotelHandler struct {
	store *db.Store // Central store providing access to all database collections
}

// Factory function to create handlers 
func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

// type HotelQueryParams struct{
// 	Rooms bool
// 	Rating int
// }
// for rooms of the specific hotel
func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	// Extract hotel ID from URL parameters
	id := c.Params("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"hotelID": oid}

	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}
// Can be extended to support query parameters for filtering
func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	// The nil filter means "get all hotels" (no conditions)
	hotels, err := h.store.Hotel.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}

	// Return hotels as JSON array
	return c.JSON(hotels)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	// Extract hotel ID from URL parameters
	id := c.Params("id")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), oid)
	if err != nil {
		return err
	}
	return c.JSON(hotel)
}
