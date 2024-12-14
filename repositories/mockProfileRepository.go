package repositories

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockProfileRepository adalah mock implementasi untuk ProfileRepository
type MockProfileRepository struct {
	mock.Mock
}

// GetDB adalah mock untuk metode GetDB yang mengembalikan objek *gorm.DB
func (m *MockProfileRepository) GetDB() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB) // Mengembalikan nilai *gorm.DB yang sesuai
}

func (m *MockProfileRepository) GetProfileByUserID(userID int) (*entities.Profil, error) {
	args := m.Called(userID)
	return args.Get(0).(*entities.Profil), args.Error(1)
}

func (m *MockProfileRepository) UpdateProfile(profile *entities.Profil) error {
	args := m.Called(profile)
	return args.Error(0)
}

// Tambahkan metode CreateProfile untuk memenuhi kontrak interface ProfileRepository
func (m *MockProfileRepository) CreateProfile(profile *entities.Profil) error {
	args := m.Called(profile)
	return args.Error(0)
}
