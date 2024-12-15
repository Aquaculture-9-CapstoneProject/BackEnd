package services

import (
	"testing"
	"time"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock untuk ArtikelRepoInterface
type MockArtikelRepo struct {
	mock.Mock
}

func (m *MockArtikelRepo) GetAll(page int, limit int) ([]entities.Artikel, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]entities.Artikel), args.Error(1)
}

func (m *MockArtikelRepo) Top3(limit int) ([]entities.Artikel, error) {
	args := m.Called(limit)
	return args.Get(0).([]entities.Artikel), args.Error(1)
}

func (m *MockArtikelRepo) FindAll(judul string, kategori string, page int, limit int) ([]entities.Artikel, error) {
	args := m.Called(judul, kategori, page, limit)
	return args.Get(0).([]entities.Artikel), args.Error(1)
}

func (m *MockArtikelRepo) FindByID(id int) (*entities.Artikel, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Artikel), args.Error(1)
}

func (m *MockArtikelRepo) Count() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func TestArtikelService(t *testing.T) {
	mockRepo := new(MockArtikelRepo)
	service := NewArtikelService(mockRepo)

	// Test GetAll
	artikel := &entities.Artikel{ID: 1, Judul: "Artikel A", Gambar: "link_gambar", CreatedAt: time.Now()}
	mockRepo.On("GetAll", 1, 9).Return([]entities.Artikel{*artikel}, nil)

	artikels, err := service.GetAll(1, 9)
	assert.NoError(t, err)
	assert.Len(t, artikels, 1)

	// Test Top3
	mockRepo.On("Top3", 3).Return([]entities.Artikel{*artikel}, nil)

	topArtikels, err := service.Top3(3)
	assert.NoError(t, err)
	assert.Len(t, topArtikels, 1)

	// Test FindAll
	mockRepo.On("FindAll", "Artikel A", "", 1, 9).Return([]entities.Artikel{*artikel}, nil)

	searchResults, err := service.FindAll("Artikel A", "", 1, 9)
	assert.NoError(t, err)
	assert.Len(t, searchResults, 1)

	// Test FindByID
	mockRepo.On("FindByID", artikel.ID).Return(artikel, nil)

	foundArtikel, err := service.FindByID(artikel.ID)
	assert.NoError(t, err)
	assert.Equal(t, artikel, foundArtikel)

	// Test Count
	mockRepo.On("Count").Return(int64(1), nil)

	count, err := service.Count()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)

	// Verifikasi bahwa semua metode mock dipanggil
	mockRepo.AssertExpectations(t)
}
