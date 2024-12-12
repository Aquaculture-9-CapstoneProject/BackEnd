package adminservices

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories/admin"
)

type AdminFilterServices interface {
	GetPaymentsByStatus(status string) ([]entities.Payment, error)
}

type adminFilterServices struct {
	repo admin.AdminFilterRepo
}

// Constructor untuk AdminPesananService
func NewAdminFilterServices(repo admin.AdminFilterRepo) AdminFilterServices {
	return &adminFilterServices{repo: repo}
}

func (s *adminFilterServices) GetPaymentsByStatus(status string) ([]entities.Payment, error) {
	return s.repo.GetPaymentsByStatus(status)
}
