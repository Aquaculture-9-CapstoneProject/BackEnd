package services_test

import (
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepoAuth struct {
	mock.Mock
}

func (m *MockRepoAuth) UserEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRepoAuth) AdminEmail(email string) (*entities.Admin, error) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Admin), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRepoAuth) DaftarAuth(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestAuthDaftarUser(t *testing.T) {
	mockRepo := new(MockRepoAuth)
	service := services.NewAuthUseCase(mockRepo)

	t.Run("succes case", func(t *testing.T) {
		mockRepo.On("UserEmail", "test@gmail.com").Return(nil, nil)
		mockRepo.On("DaftarAuth", mock.Anything).Return(nil)

		user, err := service.DaftarUser("Michael", "medan", "081378945826", "test@gmail.com", "password", "password")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "test@gmail.com", user.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("password salah", func(t *testing.T) {
		user, err := service.DaftarUser("Michael", "medan", "081378945826", "test@gmail.com", "password", "wrongpassword")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "password dan Confirmasi Password Tidak sesuai", err.Error())
	})

}
