package adminservices

import (
	"errors"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories/admin"
)

type AdminPesananServices interface {
	GetDetailedOrders() ([]map[string]interface{}, error)
}

type adminPesananServices struct {
	repopesanan admin.AdminPesananRepo
}

func NewPesananServices(repopesanan admin.AdminPesananRepo) AdminPesananServices {
	return &adminPesananServices{repopesanan: repopesanan}
}

func (ps *adminPesananServices) GetDetailedOrders() ([]map[string]interface{}, error) {
	// Panggil repository untuk mendapatkan detail pesanan
	details, err := ps.repopesanan.GetDetailedOrders()
	if err != nil {
		return nil, errors.New("gagal mendapatkan detail pesanan: " + err.Error())
	}
	return details, nil
}
