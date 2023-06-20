package models

import "time"

type Chat struct {
	ID        int              `json:"id" gorm:"primaryKey"`
	Title     string           `json:"title" gorm:"type: varchar(255)"`
	UserID    int              `json:"user_id"`
	User      UserChatResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoomID    int              `json:"room_id"`
	Room      RoomChatResponse `json:"room" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
