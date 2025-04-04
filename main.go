package main

import (
	"context"
	"flag"
	"fmt"
	"log"

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
	flag.Parse()
	app := fiber.New()
	// apiv1 := app.Group("/api/v1")
	app.Listen(*listenAddr)
}
