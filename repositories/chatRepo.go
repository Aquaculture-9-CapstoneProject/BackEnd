package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"

	"gorm.io/gorm"
)

type ChatRepoInterface interface {
	SaveChat(chat entities.Chat) error
	GetAllChat(userID int) ([]entities.Chat, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *chatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) SaveChat(chat entities.Chat) error {
	return r.db.Debug().Create(&chat).Error
}

func (r *chatRepository) GetAllChat(userID int) ([]entities.Chat, error) {
	var chats []entities.Chat
	if err := r.db.Where("user_id = ?", userID).Find(&chats).Error; err != nil {
		return nil, err
	}

	return chats, nil
}
