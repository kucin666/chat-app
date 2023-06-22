package repositories

import (
	"github.com/thxrhmn/chat-app/models"
	"gorm.io/gorm"
)

type ChatRepository interface {
	CreateChat(chat models.Chat) (models.Chat, error)
	GetChat(ID int) (models.Chat, error)
	FindChats() ([]models.Chat, error)
	FindChatsByRoomID(roomID int) ([]models.Chat, error)
	GetChatByRoomID(roomID int, chatID int) (models.Chat, error)
	UpdateChat(chat models.Chat) (models.Chat, error)
	DeleteChat(chat models.Chat, ID int) (models.Chat, error)
}

func RepositoryChat(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateChat(chat models.Chat) (models.Chat, error) {
	err := r.db.Create(&chat).Error

	return chat, err
}

func (r *repository) GetChat(ID int) (models.Chat, error) {
	var chat models.Chat
	err := r.db.First(&chat, ID).Error

	return chat, err
}

func (r *repository) FindChats() ([]models.Chat, error) {
	var chats []models.Chat
	err := r.db.Find(&chats).Error

	return chats, err
}

func (r *repository) FindChatsByRoomID(roomID int) ([]models.Chat, error) {
	var chats []models.Chat
	err := r.db.Preload("Room").Preload("User").Where("room_id=?", roomID).Find(&chats).Error

	return chats, err
}

func (r *repository) GetChatByRoomID(roomID int, chatID int) (models.Chat, error) {
	var chat models.Chat
	err := r.db.Preload("Room").Where("room_id=? AND chat_id=?", roomID, chatID).First(&chat).Error

	return chat, err
}

func (r *repository) UpdateChat(chat models.Chat) (models.Chat, error) {
	err := r.db.Save(&chat).Error

	return chat, err
}

func (r *repository) DeleteChat(chat models.Chat, ID int) (models.Chat, error) {
	err := r.db.Delete(&chat, ID).Error

	return chat, err
}
