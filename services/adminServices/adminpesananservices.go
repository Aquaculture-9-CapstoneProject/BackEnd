package adminservices

import (
	"errors"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories/admin"
)

type AdminPesananServices interface {
	GetDetailedOrders(page, perPage int) ([]map[string]interface{}, int64, error)
}

type adminPesananServices struct {
	repopesanan admin.AdminPesananRepo
}

func NewPesananServices(repopesanan admin.AdminPesananRepo) AdminPesananServices {
	return &adminPesananServices{repopesanan: repopesanan}
}

func (ps *adminPesananServices) GetDetailedOrders(page, perPage int) ([]map[string]interface{}, int64, error) {
	// Panggil repository untuk mendapatkan detail pesanan dengan pagination
	details, totalItems, err := ps.repopesanan.GetDetailedOrders(page, perPage)
	if err != nil {
		return nil, 0, errors.New("gagal mendapatkan detail pesanan: " + err.Error())
	}
	return details, totalItems, nil
}
