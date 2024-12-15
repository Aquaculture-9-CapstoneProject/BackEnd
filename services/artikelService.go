package services

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type ArtikelUseCase interface {
	GetAll(page int, limit int) ([]entities.Artikel, error)
	Top3(limit int) ([]entities.Artikel, error)
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

func (s *artikelUseCase) GetAll(page int, limit int) ([]entities.Artikel, error) {
	return s.repo.GetAll(page, limit)
}

func (s *artikelUseCase) Top3(limit int) ([]entities.Artikel, error) {
	return s.repo.Top3(limit)
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
