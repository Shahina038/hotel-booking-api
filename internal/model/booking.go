package model

import "time"

type Booking struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	RoomID     int       `json:"room_id"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateBookingRequest struct {
	RoomID   int    `json:"room_id"   validate:"required"`
	CheckIn  string `json:"check_in"  validate:"required"`
	CheckOut string `json:"check_out" validate:"required"`
}

type BookingResponse struct {
	Booking Booking `json:"booking"`
	Room    Room    `json:"room"`
}
