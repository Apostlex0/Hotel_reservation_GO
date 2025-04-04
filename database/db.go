package db
const MyHotelStore = "hotel-reservation"
const DBURI = "mongodb://localhost:27017/"
const TestDBNAME = "hotil"

type Store struct{
	User UserStore
	Hotel HotelStore
	Room RoomStore
	Booking BookingStore
}