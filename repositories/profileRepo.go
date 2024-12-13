package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfileByUserID(userID int) (*entities.Profil, error)
	CreateProfile(profile *entities.Profil) error
	UpdateProfile(profile *entities.Profil) error
	GetDB() *gorm.DB
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *profileRepository) CreateProfile(profile *entities.Profil) error {
	return r.db.Create(profile).Error
}

func (r *profileRepository) GetProfileByUserID(userID int) (*entities.Profil, error) {
	var profile entities.Profil
	err := r.db.Preload("User").Where("user_id = ?", userID).First(&profile).Error
	return &profile, err
}

func (r *profileRepository) UpdateProfile(profile *entities.Profil) error {
	return r.db.Save(profile).Error
}
