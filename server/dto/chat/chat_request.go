package chatdto

import "time"

type ChatRequest struct {
	Title     string    `json:"title" form:"title" validate:"required"`
	UserID    int       `json:"user_id"`
	RoomID    int       `json:"room_id" form:"room_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
