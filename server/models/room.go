package models

import "time"

type Room struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	CreatedBy int       `json:"user_id" gorm:"type: int"`
	CreatedAt time.Time `json:"created_at"`
	UpadtedAt time.Time `json:"updated_at"`
}

type RoomChatResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (RoomChatResponse) TableName() string {
	return "rooms"
}
