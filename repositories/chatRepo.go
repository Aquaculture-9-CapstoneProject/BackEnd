package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"

	"gorm.io/gorm"
)

type ChatRepoInterface interface {
	SaveChat(chat entities.Chat) (entities.Chat, error)
	GetAllChat(userID int) ([]entities.Chat, error)
	GetChatByID(chatID int) (entities.Chat, error)
	GetRecommendedProducts() ([]entities.Product, error)
	GetProductByName(productName string) (entities.Product, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *chatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) SaveChat(chat entities.Chat) (entities.Chat, error) {
	if err := r.db.Debug().Preload("user").Create(&chat).Error; err != nil {
		return entities.Chat{}, err
	}
	return chat, nil
}

func (r *chatRepository) GetAllChat(userID int) ([]entities.Chat, error) {
	var chats []entities.Chat
	if err := r.db.Preload("User").Where("user_id = ?", userID).Find(&chats).Error; err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *chatRepository) GetChatByID(chatID int) (entities.Chat, error) {
	var chat entities.Chat
	if err := r.db.Preload("User").First(&chat, chatID).Error; err != nil {
		return entities.Chat{}, err
	}
	return chat, nil
}

func (repo *chatRepository) GetRecommendedProducts() ([]entities.Product, error) {
	var products []entities.Product
	err := repo.db.Where("rating >= ?", 3.0).Order("rating desc").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *chatRepository) GetProductByName(productName string) (entities.Product, error) {
	// Cari produk berdasarkan nama dalam database
	var product entities.Product
	err := repo.db.Where("nama LIKE ?", "%"+productName+"%").First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
