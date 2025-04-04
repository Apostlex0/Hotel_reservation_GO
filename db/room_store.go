package db

import (
	"context"

	"github.com/Apostlex0/Hotel_reservation_GO/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
// Any implementation of RoomStore must provide these methods
type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error) // Add a new room
	GetRooms(context.Context, bson.M) ([]*types.Room, error)      // Get rooms with optional filters
}

type MongoRoomStore struct {
	client     *mongo.Client     
	coll       *mongo.Collection 
	HotelStore                   
}
func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		coll:       client.Database(MyHotelStore).Collection("rooms"),
		HotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	// Insert the room document
	resp, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)

	// Create filter and update to add the room ID to the hotel's rooms array
	filter := bson.M{"_id": room.HotelID}
	update := bson.M{"$push": bson.M{"rooms": room.ID}}

	// Update the hotel document to include this room
	if err := s.HotelStore.Update(ctx, filter, update); err != nil {
		return nil, err
	}

	return room, nil
}

func (s *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error) {
	// Find rooms matching the filter
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Decode all results into the rooms slice
	var rooms []*types.Room
	if err := resp.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}
