package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type ArtikelUseCase interface {
	Create(artikel *entities.Artikel) (*entities.Artikel, error)
	Update(artikel *entities.Artikel) (*entities.Artikel, error)
	Delete(id int) error
	GetAll(page int, limit int) ([]entities.Artikel, error)
	FindAll(judul string, kategori string, page int, limit int) ([]entities.Artikel, error)
	FindByID(id int) (*entities.Artikel, error)
	Count() (int64, error)
}

type artikelUseCase struct {
	repo repositories.ArtikelRepoInterface
}

func NewArtikelService(repo repositories.ArtikelRepoInterface) *artikelUseCase {
	return &artikelUseCase{repo: repo}
}

func (s *artikelUseCase) Create(artikel *entities.Artikel) (*entities.Artikel, error) {
	return s.repo.Create(artikel)
}

func (s *artikelUseCase) Update(artikel *entities.Artikel) (*entities.Artikel, error) {
	return s.repo.Update(artikel)
}

func (s *artikelUseCase) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *artikelUseCase) GetAll(page int, limit int) ([]entities.Artikel, error) {
	return s.repo.GetAll(page, limit)
}

func (s *artikelUseCase) FindAll(judul string, kategori string, page int, limit int) ([]entities.Artikel, error) {
	artikels, err := s.repo.FindAll(judul, kategori, page, limit)
	if err != nil {
		return nil, err
	}
	return artikels, nil
}

func (s *artikelUseCase) FindByID(id int) (*entities.Artikel, error) {
	return s.repo.FindByID(id)
}

func (s *artikelUseCase) Count() (int64, error) {
	return s.repo.Count()
}
