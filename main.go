package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/Apostlex0/Hotel_reservation_GO/api"
	"github.com/Apostlex0/Hotel_reservation_GO/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"

func main() {
	listenAddr := flag.String("listenAddr", ":5001", "The listen address of the API server")
	flag.Parse()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)

	userStore := db.NewMongoUserStore(client)
	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)
	bookingStore := db.NewMongoBookingStore(client)

	store := &db.Store{
		Hotel:   hotelStore,
		User:    userStore,
		Room:    roomStore,
		Booking: bookingStore,
	}
	userHandler := api.NewUserHandler(userStore)
	hotelHandler := api.NewHotelHandler(store)
	roomHandler := api.NewRoomHandler(store)

	app := fiber.New()

	apiv1 := app.Group("/api/v1")
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Put("user/:id", userHandler.HandlePutUser)

	//hotel routes here
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)    // Get all hotels
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel) // Get a specific hotel
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)
	apiv1.Get("/room",roomHandler.HandleGetRooms)

	apiv1.Post("/room/:id/book",roomHandler.HandleBookRoom)

	app.Listen(*listenAddr)
}
