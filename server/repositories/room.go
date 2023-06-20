package repositories

import (
	"github.com/thxrhmn/chat-app/models"
	"gorm.io/gorm"
)

type RoomRepository interface {
	CreateRoom(room models.Room) (models.Room, error)
	GetRoom(ID int) (models.Room, error)
	FindRooms() ([]models.Room, error)
	UpdateRoom(room models.Room) (models.Room, error)
	DeleteRoom(room models.Room, ID int) (models.Room, error)
}

func RepositoryRoom(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateRoom(room models.Room) (models.Room, error) {
	err := r.db.Create(&room).Error

	return room, err
}

func (r *repository) GetRoom(ID int) (models.Room, error) {
	var room models.Room
	err := r.db.First(&room, ID).Error

	return room, err
}

func (r *repository) FindRooms() ([]models.Room, error) {
	var rooms []models.Room
	err := r.db.Find(&rooms).Error

	return rooms, err
}

func (r *repository) UpdateRoom(room models.Room) (models.Room, error) {
	err := r.db.Save(&room).Error

	return room, err
}

func (r *repository) DeleteRoom(room models.Room, ID int) (models.Room, error) {
	err := r.db.Delete(&room, ID).Error

	return room, err
}
