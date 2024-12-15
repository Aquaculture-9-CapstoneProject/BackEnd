package adminservices

import (
	"testing"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock untuk AdminArtikelRepoInterface
type MockAdminArtikelRepo struct {
	mock.Mock
}

func (m *MockAdminArtikelRepo) Create(artikel *entities.Artikel) (*entities.Artikel, error) {
	args := m.Called(artikel)
	return args.Get(0).(*entities.Artikel), args.Error(1)
}

func (m *MockAdminArtikelRepo) Update(artikel *entities.Artikel) (*entities.Artikel, error) {
	args := m.Called(artikel)
	return args.Get(0).(*entities.Artikel), args.Error(1)
}

func (m *MockAdminArtikelRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAdminArtikelRepo) GetAll(page int, limit int) ([]entities.Artikel, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]entities.Artikel), args.Error(1)
}

func (m *MockAdminArtikelRepo) FindAll(judul string, kategori string, page int, limit int) ([]entities.Artikel, error) {
	args := m.Called(judul, kategori, page, limit)
	return args.Get(0).([]entities.Artikel), args.Error(1)
}

func (m *MockAdminArtikelRepo) FindByID(id int) (*entities.Artikel, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Artikel), args.Error(1)
}

func (m *MockAdminArtikelRepo) Count() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func TestAdminArtikelService(t *testing.T) {
	mockRepo := new(MockAdminArtikelRepo)
	service := NewAdminArtikelService(mockRepo)

	// Test Create
	artikel := &entities.Artikel{ID: 1, Judul: "Artikel A", Gambar: "link_gambar"}
	mockRepo.On("Create", artikel).Return(artikel, nil)

	createdArtikel, err := service.Create(artikel)
	assert.NoError(t, err)
	assert.Equal(t, artikel, createdArtikel)

	// Test Update
	artikel.Judul = "Artikel A Updated"
	mockRepo.On("Update", artikel).Return(artikel, nil)

	updatedArtikel, err := service.Update(artikel)
	assert.NoError(t, err)
	assert.Equal(t, artikel, updatedArtikel)

	// Test Delete
	mockRepo.On("Delete", artikel.ID).Return(nil)

	err = service.Delete(artikel.ID)
	assert.NoError(t, err)

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

	// Test GetAll
	mockRepo.On("GetAll", 1, 10).Return([]entities.Artikel{*artikel}, nil)

	artikels, err := service.GetAll(1, 10)
	assert.NoError(t, err)
	assert.Len(t, artikels, 1)

	// Test FindAll
	mockRepo.On("FindAll", "Artikel A", "", 1, 10).Return([]entities.Artikel{*artikel}, nil)

	searchResults, err := service.FindAll("Artikel A", "", 1, 10)
	assert.NoError(t, err)
	assert.Len(t, searchResults, 1)

	// Verifikasi bahwa semua metode mock dipanggil
	mockRepo.AssertExpectations(t)
}
