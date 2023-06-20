package room

import (
	"time"

	"github.com/thxrhmn/chat-app/models"
)

type RoomResponse struct {
	Name          string                  `json:"name"`
	CreatedById   int                     `json:"created_by_id"`
	CreatedByUser models.UserRoomResponse `json:"created_by_user"`
	CreatedAt     time.Time               `json:"created_at"`
}
