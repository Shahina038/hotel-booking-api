package model

import "time"

type Room struct {
	ID             int       `json:"id"`
	RoomNumber     string    `json:"room_number"`
	Type           string    `json:"type"`
	Description    string    `json:"description"`
	PricePerNight  float64   `json:"price_per_night"`
	Capacity       int       `json:"capacity"`
	IsAvailable    bool      `json:"is_available"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateRoomRequest struct {
	RoomNumber    string  `json:"room_number"    validate:"required"`
	Type          string  `json:"type"           validate:"required,oneof=single double suite deluxe"`
	Description   string  `json:"description"`
	PricePerNight float64 `json:"price_per_night" validate:"required,gt=0"`
	Capacity      int     `json:"capacity"        validate:"required,gt=0"`
}

type UpdateRoomRequest struct {
	Type          string  `json:"type"            validate:"omitempty,oneof=single double suite deluxe"`
	Description   string  `json:"description"`
	PricePerNight float64 `json:"price_per_night" validate:"omitempty,gt=0"`
	Capacity      int     `json:"capacity"        validate:"omitempty,gt=0"`
	IsAvailable   *bool   `json:"is_available"`
}

type RoomSearchRequest struct {
	Type      string  `json:"type"`
	MinPrice  float64 `json:"min_price"`
	MaxPrice  float64 `json:"max_price"`
	CheckIn   string  `json:"check_in"`
	CheckOut  string  `json:"check_out"`
	Capacity  int     `json:"capacity"`
}
