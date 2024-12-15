package adminservices

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories/admin"
)

type AdminArtikelUseCase interface {
	Create(artikel *entities.Artikel) (*entities.Artikel, error)
	Update(artikel *entities.Artikel) (*entities.Artikel, error)
	Delete(id int) error
	GetAll(page int, limit int) ([]entities.Artikel, error)
	FindAll(judul string, kategori string, page int, limit int) ([]entities.Artikel, error)
	FindByID(id int) (*entities.Artikel, error)
	Count() (int64, error)
}

type adminArtikelUseCase struct {
	repo admin.AdminArtikelRepoInterface
}

func NewAdminArtikelService(repo admin.AdminArtikelRepoInterface) *adminArtikelUseCase {
	return &adminArtikelUseCase{repo: repo}
}

func (s *adminArtikelUseCase) Create(artikel *entities.Artikel) (*entities.Artikel, error) {
	return s.repo.Create(artikel)
}

func (s *adminArtikelUseCase) Update(artikel *entities.Artikel) (*entities.Artikel, error) {
	return s.repo.Update(artikel)
}

func (s *adminArtikelUseCase) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *adminArtikelUseCase) GetAll(page int, limit int) ([]entities.Artikel, error) {
	return s.repo.GetAll(page, limit)
}

func (s *adminArtikelUseCase) FindAll(judul string, kategori string, page int, limit int) ([]entities.Artikel, error) {
	artikels, err := s.repo.FindAll(judul, kategori, page, limit)
	if err != nil {
		return nil, err
	}
	return artikels, nil
}

func (s *adminArtikelUseCase) FindByID(id int) (*entities.Artikel, error) {
	return s.repo.FindByID(id)
}

func (s *adminArtikelUseCase) Count() (int64, error) {
	return s.repo.Count()
}
