package adminservices

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories/admin"
)

type AdminFilterServices interface {
	GetPaymentsByStatus(status string) ([]entities.Payment, error)
	GetPaymentsByStatusBarang(statusBarang string) ([]entities.Payment, error)
}

type adminFilterServices struct {
	repo admin.AdminFilterRepo
}

func NewAdminFilterServices(repo admin.AdminFilterRepo) AdminFilterServices {
	return &adminFilterServices{repo: repo}
}

func (s *adminFilterServices) GetPaymentsByStatus(status string) ([]entities.Payment, error) {
	return s.repo.GetPaymentsByStatus(status)
}

func (s *adminFilterServices) GetPaymentsByStatusBarang(statusBarang string) ([]entities.Payment, error) {
	return s.repo.GetPaymentsByStatusBarang(statusBarang)
}
