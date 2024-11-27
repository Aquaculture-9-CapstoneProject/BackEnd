package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"gorm.io/gorm"
)

type RepoAuth interface {
	DaftarAuth(user *entities.User) error
	UserEmail(email string) (*entities.User, error)
	AdminEmail(email string) (*entities.Admin, error)
}

type repoAuth struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) RepoAuth {
	return &repoAuth{db: db}
}

func (r *repoAuth) DaftarAuth(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *repoAuth) UserEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ? ", email).First(&user).Error
	return &user, err
}

func (r *repoAuth) AdminEmail(email string) (*entities.Admin, error) {
	var admin entities.Admin
	err := r.db.Where("email = ? ", email).First(&admin).Error
	return &admin, err
}
