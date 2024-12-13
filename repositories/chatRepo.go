package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"

	"gorm.io/gorm"
)

type ChatRepoInterface interface {
	SaveChat(chat entities.Chat) (entities.Chat, error)
	GetAllChat(userID int) ([]entities.Chat, error)
	GetChatByID(chatID int) (entities.Chat, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *chatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) SaveChat(chat entities.Chat) (entities.Chat, error) {
	if err := r.db.Debug().Create(&chat).Error; err != nil {
		return entities.Chat{}, err
	}
	return chat, nil
}

func (r *chatRepository) GetAllChat(userID int) ([]entities.Chat, error) {
	var chats []entities.Chat
	if err := r.db.Where("user_id = ?", userID).Find(&chats).Error; err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *chatRepository) GetChatByID(chatID int) (entities.Chat, error) {
	var chat entities.Chat
	if err := r.db.First(&chat, chatID).Error; err != nil {
		return entities.Chat{}, err
	}
	return chat, nil
}
