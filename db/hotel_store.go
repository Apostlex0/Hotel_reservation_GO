package db

import (
	"context"

	"github.com/Apostlex0/Hotel_reservation_GO/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Any implementation of HotelStore must provide these methods
type HotelStore interface {
	Insert(context.Context, *types.Hotel) (*types.Hotel, error)             // Add a new hotel
	Update(context.Context, bson.M, bson.M) error                           // Update hotel information
	GetHotels(context.Context, bson.M) ([]*types.Hotel, error)              // Get hotels with optional filters
	GetHotelByID(context.Context, primitive.ObjectID) (*types.Hotel, error) // Find a hotel by ID
}


type MongoHotelStore struct {
	client *mongo.Client     // MongoDB client connection
	coll   *mongo.Collection // Reference to the hotels collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(MyHotelStore).Collection("hotels"),
	}
}

func (s *MongoHotelStore) Insert(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	// Insert the hotel document
	resp, err := s.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	// Update the hotel object with the generated ID
	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) Update(ctx context.Context, filter, update bson.M) error {
	// Update the hotel document
	_, err := s.coll.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*types.Hotel, error) {
	// Find hotels matching the filter
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Decode all results into the hotels slice
	var hotels []*types.Hotel
	if err := resp.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}
func (s *MongoHotelStore) GetHotelByID(ctx context.Context, id primitive.ObjectID) (*types.Hotel, error) {
	var hotel types.Hotel

	// Find and decode the hotel document
	if err := s.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&hotel); err != nil {
		return nil, err
	}
	return &hotel, nil
}
