package models

import "time"

type Room struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	CreatedBy int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"type: varchar(255)"`
	UpadtedAt time.Time `json:"updated_at" gorm:"type: varchar(255)"`
}
