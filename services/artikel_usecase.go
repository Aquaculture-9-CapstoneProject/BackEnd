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
	FindByID(id int) (*entities.Artikel, error)
	GetAdminByID(id int) (*entities.Admin, error)
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
	return s.repo.FindAll(page, limit)
}

func (s *artikelUseCase) FindByID(id int) (*entities.Artikel, error) {
	return s.repo.FindByID(id)
}

func (s *artikelUseCase) GetAdminByID(id int) (*entities.Admin, error) {
	return s.repo.GetAdminByID(id)
}

func (s *artikelUseCase) Count() (int64, error) {
	return s.repo.Count()
}