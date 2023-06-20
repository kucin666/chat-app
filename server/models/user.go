package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name" gorm:"type:varchar(255)"`
	Username     string    `json:"username" gorm:"unique;not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	Password     string    `json:"password" gorm:"type:varchar(255)"`
	IsAdmin      bool      `json:"is_admin" gorm:"type:boolean"`
	ProfileImage string    `json:"profile_image" gorm:"type:varchar(255)"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type UserRoomResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserChatResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func (UserRoomResponse) TableName() string {
	return "users"
}

func (UserChatResponse) TableName() string {
	return "users"
}
