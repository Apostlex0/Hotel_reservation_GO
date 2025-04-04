package db

import (
	"context"
	"fmt"

	"github.com/Apostlex0/Hotel_reservation_GO/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Name of the MongoDB collection for users
const usesrColl = "users"

// UserStore defines the interface for user data operations
type UserStore interface {
	GetUserByEmail(context.Context, string) (*types.User, error)        
	GetUserById(context.Context, string) (*types.User, error)           
	GetUsers(context.Context) ([]*types.User, error)                    
	InsertUser(context.Context, *types.User) (*types.User, error)       
	DeleteUser(context.Context, string) error                           
	UpdateUser(ctx context.Context, filter bson.M, update bson.M) error // Update user information
	Drop(context.Context) error                                         
}

type MongoUserStore struct {
	client *mongo.Client     
	MyHotelStore string            
	coll   *mongo.Collection 
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(MyHotelStore).Collection(usesrColl),
	}
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	var user types.User

	oid, err := primitive.ObjectIDFromHex(id) //string id to ObjectID
	if err != nil {
		return nil, err
	}

	// Find the user document and decode it into the user variable
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}


func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User

	// Find all user documents
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	// Insert the user document
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// Update the user object with the generated ID
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

// DeleteUser removes a user from the database
func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	// Convert string ID to MongoDB ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Delete the user document
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, update bson.M) error {

	_, err := s.coll.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("----dropping")
	return s.coll.Drop(ctx)
}

func (s *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	// Find and decode the user document
	if err := s.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
