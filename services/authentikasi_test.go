package services_test

import (
	"errors"
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

	t.Run("email sudah ada", func(t *testing.T) {
		mockRepo.On("UserEmail", "test123@gmail.com").Return(&entities.User{Email: "test123@gmail.com"}, nil)

		user, err := service.DaftarUser("John Doe", "Address", "123456789", "test123@gmail.com", "password", "password")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "email sudah digunakan", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestAuthloginUser(t *testing.T) {
	mockRepo := new(MockRepoAuth)
	service := services.NewAuthUseCase(mockRepo)

	t.Run("succes case", func(t *testing.T) {
		mockRepo.On("UserEmail", "test123@gmail.com").Return(&entities.User{Email: "test123@gmail.com", Password: "password"}, nil)

		user, err := service.LoginUser("test123@gmail.com", "password")

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "test123@gmail.com", user.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("password salah", func(t *testing.T) {
		mockRepo.On("UserEmail", "test123@gmail.com").Return(&entities.User{Email: "test123@gmail.com", Password: "password"}, nil)

		user, err := service.LoginUser("test123@gmail.com", "wrongpassword")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "password salah", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo.On("UserEmail", "test1234@gmail.com").Return(nil, errors.New("user not found"))

		user, err := service.LoginUser("test1234@gmail.com", "password")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "user not found", err.Error())
		mockRepo.AssertExpectations(t)
	})

}

func TestAuthloginAdmin(t *testing.T) {
	mockRepo := new(MockRepoAuth)
	service := services.NewAuthUseCase(mockRepo)

	t.Run("succes case", func(t *testing.T) {
		mockRepo.On("AdminEmail", "admin123@gmail.com").Return(&entities.Admin{Email: "admin123@gmail.com", Password: "adminipassword"}, nil)

		admin, err := service.LoginAdmin("admin123@gmail.com", "adminipassword")

		assert.NoError(t, err)
		assert.NotNil(t, admin)
		assert.Equal(t, "admin123@gmail.com", admin.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("password salah", func(t *testing.T) {
		mockRepo.On("AdminEmail", "admin123@gmail.com").Return(&entities.Admin{Email: "admin123@gmail.com", Password: "adminipassword"}, nil)

		admin, err := service.LoginAdmin("admin123@gmail.com", "passwordsalah")

		assert.Error(t, err)
		assert.Nil(t, admin)
		assert.Equal(t, "password salah", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("admin tidak ada", func(t *testing.T) {
		mockRepo.On("AdminEmail", "admin1234@gmail.com").Return(nil, errors.New("admin not found"))

		admin, err := service.LoginAdmin("admin1234@gmail.com", "adminipassword")

		assert.Error(t, err)
		assert.Nil(t, admin)
		assert.Equal(t, "admin not found", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
