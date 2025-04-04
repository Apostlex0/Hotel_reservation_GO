package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// RoomType defines the available types of rooms in the hotel
// Using iota for auto-incrementing integer constants
type RoomType int

const(
	_ RoomType = iota      // Skip the first value (0)
	SingleRoomType         // Value: 1 - A room for one person
	DoubleRoomType         // Value: 2 - A room for two people
	SeaSideRoomType        // Value: 3 - A room with a sea view
	DeluxRoomType          // Value: 4 - A luxury room with premium amenities
)

// Hotel represents a hotel in the reservation system
// Contains basic information about the hotel and references to its rooms
type Hotel struct{
	ID 		 primitive.ObjectID     `bson:"_id,omitempty" json:"id"`       
	Name 	 string 		        `bson:"name" json:"name"`               
	Location string 	            `bson:"location" json:"location"`      
	Rooms 	 []primitive.ObjectID	`bson:"rooms" json:"rooms"`             
	Rating 	 int					`bson:"rating" json:"rating"`           
}

// Room represents an individual room in a hotel
// Contains details about the room's features and pricing
type Room struct{
	ID        primitive.ObjectID     `bson:"_id,omitempty" json:"id,omitempty"` 
	Seaside	  bool 				     `bson:"seaside" json:"seaside"`           
	Size 	  string			     `bson: "size" json:"size"`               
	Price 	  float64			     `bson:"price",json:"price"`              
	HotelID   primitive.ObjectID     `bson:"hotelID" json:"hotelID"`          
}