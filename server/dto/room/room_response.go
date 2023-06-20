package roomdto

import (
	"time"

	"github.com/thxrhmn/chat-app/models"
)

type RoomResponse struct {
	ID            int                     `json:"id"`
	Name          string                  `json:"name"`
	CreatedById   int                     `json:"created_by_id"`
	CreatedByUser models.UserRoomResponse `json:"created_by_user"`
	CreatedAt     time.Time               `json:"created_at"`
}
